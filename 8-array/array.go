package main

import "fmt"

func main() {

	// 初始化 int 为 0
	var a [5]int
	fmt.Println("emp:", a)

	// 普通的数组赋值和取值
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// 数组的长度和容量都是一样的
	fmt.Println("len:", len(a))
	fmt.Println("cap:", cap(a))

	// 另外一种赋值方式，直接在声明的时候赋值
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	// 定义二维数组，然后使用 for 循环来赋值
	var twoD [2][3]int
	// 以上已经给数组分配空间了
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
}
