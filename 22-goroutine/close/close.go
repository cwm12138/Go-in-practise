package main

import "fmt"

func main() {

	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, status := <-jobs
			// 第一个返回值是 chan 的值，第二个返回值是是否还有值
			fmt.Println(status)
			if status {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for i := 1; i <= 3; i++ {
		jobs <- i
		fmt.Println("send job: ", i)
	}

	close(jobs)
	// 如果不关闭channel 当channel 没有值的时候，那么就会发生死锁
	// TODO 研究死锁问题
	// 经过这些例子发现 channel 如果没有值去读 会发生死锁，如果关闭了，再往里写也会发生死锁
	fmt.Println("send all jobs")

	<-done
}
