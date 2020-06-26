package main

import (
	"fmt"
	"time"
)

// 参数是一个 Bool 类型的 channel
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// 往 channel 中加入信号 通知该 func 已执行完毕
	done <- true
}

func main() {

	done := make(chan bool, 1)
	go worker(done)

	// 一个 channel 会 block 函数返回
	// 直到这个channel有值，那么取消 block 程序继续运行
	// 翻译过来就是，直到 worker 执行完毕之前程序不会退出，不然程序直接退出，看不到 goroutine 的执行效果
	<-done

}
