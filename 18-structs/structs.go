package main

import "fmt"

type person struct {
	name string
	age  int
}

// 实现类似构造函数的功能
// return 一个数组
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 22
	return &p
}

// 这个结构体的方法
func (p *person) talk() {
	fmt.Println(p.name + " say hello")
}

func main() {
	fmt.Println(person{name: "foo", age: 23})
	fmt.Println(person{"bar", 22})
	fmt.Println(person{name: "cwm"})
	fmt.Println(&person{name: "ann", age: 12})

	fmt.Println(newPerson("mike"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 50
	fmt.Println(s, sp)

	spNew := new(person)
	fmt.Println(spNew)

}
