package main

import (
	"fmt"
	"math"
)

// go by example 提供了一个类似于继承的例子
// http://www.golang-book.com/books/intro 这个书中的 interface 章节提供了接口作为参数的例子，很受启发

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	wdith, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.wdith * r.height
}

func (r rect) perim() float64 {
	return (r.wdith + r.height) * 2
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())

}

func main() {

	// 声明继承 interface 的 struct
	r := rect{10, 10}
	c := circle{4}

	// 因为是实现的 interface 所以这个实例可以作为参数
	measure(r)
	measure(c)

}
