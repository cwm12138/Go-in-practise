package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}

}

func main() {

	f("direct")

	// 采用 go 这个关键词启动一个 goroutine 这是一个轻量级的并发操作
	go f("goroutine")
	fmt.Println("hello")

	go func(msg string) {
		fmt.Println(msg)
	}("msg")

	// 如果不 sleep 主程序 goroutine 结束之后其他协程也会退出，不会看到输出
	time.Sleep(time.Second)
	fmt.Println("done")
}
