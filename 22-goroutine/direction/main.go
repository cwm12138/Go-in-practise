package main

import "fmt"

// 通过指定 channel 的方向可以增加程序的类型安全
// chan<- 表示这是一个接受者
// <-chan 表示这是一个发送者
func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg

}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "Test message")
	pong(pings, pongs)

	fmt.Println(<-pongs)
}
