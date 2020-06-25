package main

import "fmt"

// 当有多个返回值的时候记得用括号
func vals() (int, int) {
	return 3, 7
}

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// 不需要的返回值依然使用 blank identifier _ 去忽视它
	_, c := vals()
	fmt.Println(c)

}
