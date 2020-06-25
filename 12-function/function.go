package main

import "fmt"

// 注意的地方就是定义变量和返回值都是在后面，这点和 java 不同
func plus(a int, b int) int {

	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plus(1, 2)
	fmt.Println(res)
	res = plusPlus(1, 2, 3)
	fmt.Println("1 + 2 + 3 = ", res)
}
