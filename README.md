# golang dispatcher
golang 调度器
---
之前写代码的时候，随意开goroutine，写完项目后就遇到两个问题。
- 1、goroutine的数量在某些时候太多，严重影响性能
- 2、当程序关闭的时候也无法保证goroutine的任务都已经完成

于是就写了这个调度器来解决这些问题。

---

