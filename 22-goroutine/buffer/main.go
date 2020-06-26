package main

import "fmt"

func bufferChannel() {
	bufferMsg := make(chan string, 2)

	bufferMsg <- "1"
	bufferMsg <- "2"

	// 只能装入数量限制的长度
	fmt.Println(<-bufferMsg)
	bufferMsg <- "3"
	fmt.Println(<-bufferMsg)
}

// 默认是不带缓冲的 channel
func main() {
	// 而且发现 buffer channel 可以在非协程中使用
	// buffer := make(chan string)

	// buffer <- "test1"
	// msg := <-buffer
	// // buffer <- "test2"

	// fmt.Println(msg)
	// fmt.Println(msg)

	bufferChannel()
}
