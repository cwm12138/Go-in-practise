package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	// 作为参数传递是需要加上 ... 的
	nums := []int{1, 2, 3, 4, 5, 6}
	sum(nums...)

}
