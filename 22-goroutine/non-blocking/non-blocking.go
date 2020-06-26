package main

import "fmt"

// 如果是 buffer chan 自动就是non blocking 的
func main() {
	message := make(chan string)
	signals := make(chan bool)

	// 接收， chan 用于 发送者
	select {
	case msg := <-message:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	// TODO make it work
	select {
	case message <- msg:
		fmt.Println("send message", msg)
	default:
		fmt.Println("no messgae sent")
	}

	select {
	case msg := <-message:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("reseived signal", sig)
	default:
		fmt.Println("no activity")
	}

}
