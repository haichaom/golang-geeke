## Question

1. 基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。

## Output

PS C:\Users\Administrator\Documents\golang-geeke\excercise\errorgroup> go run main.go<br>
main:: start http Server &{:80 <nil> <nil> 0s 0s 0s 0s 0 map[] <nil> <nil> <nil> <nil> 0 0 {0 {0 0}} <nil> {0 0} map[] map[] <nil> []}<br>
func-helloWorld:: called<br>
main-select:: Signal received:  interrupt<br>
main-select:: ctx Done<br>
main:: Shut down httpServer<br>
main:: Found error:  context canceled<br>


