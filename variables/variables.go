package main

import "fmt"

func main() {

	// declares 1 or more variables
	var a = "inital"
	fmt.Println(a)

	// declare multiple variables at once
	var b, c int = 1, 2
	fmt.Println(b, c)

	// 自动类型推断
	var d = true
	fmt.Println(d)

	// 可以不用初始化 自动初始化
	var e int
	fmt.Println(e)

	// 隐式声明
	// var f string = "apple"
	f := "apple"
	fmt.Println(f)

}
