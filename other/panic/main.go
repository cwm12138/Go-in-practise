package main

import "fmt"

func main() {
	// defer 无论程序是否 panic 是否出错都会运行的
	// 前面是一个匿名函数，花括号之后调用他 没毛病！
	defer func() {
		str := recover()
		fmt.Println(str)
	}()

	panic("This is PANIC, mean's this program is broken")

}
