package main

import "fmt"

func main() {
	// make 一个 string 类型的 channel
	messages := make(chan string)

	// store <- input
	go func() {
		messages <- "hello This is channel"
	}()

	// 这里好像会等管道把消息发送过来才会执行，不然就会一直等待
	// This property allowed us to wait at the end of our program for the "ping" message without having to use any other synchronization.
	msg := <-messages

	fmt.Println(msg)
}
