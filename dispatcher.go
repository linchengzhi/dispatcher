package dispatcher

import (
	"sync"
)

const (
	Queue_Len = 10000
	Work_Len = 1000
)

type Interface interface {
	Do(data interface{})
}

type Dispatcher struct {
	Interface
	queue   chan interface{} //数据队列
	workLen int //协程数
	exit    chan bool //退出协程
	close   bool
	wg      *sync.WaitGroup	//用于保证所有的任务都已完成才关闭调度器
}

//新建一个调度器
func New(queueLen int, workLen int, do Interface) *Dispatcher {
	d := new(Dispatcher)
	if queueLen == 0 {
		queueLen = Queue_Len
	}
	if workLen == 0 {
		workLen = Work_Len
	}
	d.queue = make(chan interface{}, queueLen)
	d.workLen = workLen
	d.Interface = do
	d.exit = make(chan bool)
	d.wg = new(sync.WaitGroup)

	for i := 0; i < d.workLen; i++ {
		go d.start()
	}
	return d
}
//启动协程
func (d *Dispatcher) start() {
	for {
		select {
		case j := <-d.queue:
			d.Interface.Do(j)
			d.wg.Done()
		case <-d.exit:
			return
		}
	}
}

//把任务添加到队列
func (d *Dispatcher) Add(data interface{}) {
	if d.close {
		return
	}
	d.queue <- data
	d.wg.Add(1)
}

//优雅退出调度器
func (d *Dispatcher) Exit() {
	d.wg.Wait()
	for i := 0; i < d.workLen; i++ {
		d.exit <- true
	}
	d.close = true
}