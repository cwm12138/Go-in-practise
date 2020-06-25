package main

import "fmt"

type rect struct {
	width, height int
}

// 成员方法
func (r *rect) area() int {
	return r.width * r.height
}

// 可是是一个指针 也可以是一个值
// 传指针 可以避免在方法中复制一遍值
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 10}

	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim: ", rp.perim())

}
