# 学习笔记
## 作业
基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够 一个退出，全部注销退出。

## 笔记📒
1、
2、nerver start a goroutine without knowning when it will stop。
要对每个goroutine的生命周期进行掌控。

1、要注意要谨慎在goroutine里面使用log.Fatal(err)，
log.Fatal底层调用os.exit，一般在main函数或者init函数使用。

sync.atomic
Copy on Write（写时复制）

errgroup

context是线程安全的

谁调用谁

reference
https://github.com/bin10203/Go-000/tree/main/Week03