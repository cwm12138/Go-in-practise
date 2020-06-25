package main

import "fmt"

// key-value
func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map", m)

	// 获取值有两个返回值，第一个是取到的值是多少，第二个是是否存在
	v1, v2 := m["k1"]
	fmt.Println(v1, v2)

	// 内置一个删除函数 参数1 待删除 map 参数2 k2 key
	delete(m, "k2")
	t1, t2 := m["k2"]
	// 如果这里把 key 删掉，那么会提示不存在并且返回初始值
	fmt.Println(t1, t2)

	// 当然 也可以在声明的时候赋值
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(n)
}
