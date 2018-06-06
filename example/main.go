package main

import (
	"github.com/linchengzhi/dispatcher"
	"fmt"
	"time"
)

type work struct {
}

func main() {
	w := new(work)
	d := dispatcher.New(100, 10, w)
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			d.Add(123)
		} else {
			d.Add("hello world")
		}
	}
	fmt.Println("all add")
	d.Exit()
	fmt.Println("dispatcher exit")
	time.Sleep(20 * time.Second)
}

func (this *work)Do(data interface{}) {
	time.Sleep(1*time.Second)
	switch data.(type) {
	case int:
		fmt.Println(data.(int) + 2210)
	case string:
		fmt.Println("test:", data.(string))
	default:
		fmt.Printf("other type data = %v /n", data)
	}
}