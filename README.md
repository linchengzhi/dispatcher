# golang goroutine dispatcher
golang 协程调度器
---
之前写代码的时候，随意开goroutine，写完项目后就遇到两个问题。
- 1、goroutine的数量在某些时候太多，严重影响性能
- 2、当程序关闭的时候也无法保证goroutine的任务都已经完成

于是就写了这个调度器来解决这些问题。调度器可保证协程控制在最大并发数内，并发数和服务器性能有关

---
说明：
```go
    w := new(work) //实现Interface.Do接口 调度器会调用此接口做任务
    参数 队列长度， 最大并发数， Interface.Do接口
	d := dispatcher.New(100, 10, w) //新建一个调度器
	d.Add(123) //把任务加入调度器
	d.Exit() //退出调度器
	//详细请参考example
```

