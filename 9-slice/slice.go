package main

import "fmt"

func main() {
	// 初始化一个 slice
	s := make([]string, 3)
	fmt.Println(s)
	//输出结果 为 zero-value

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// 可以往 slice 中添加新元素
	// appen() 第一个参数的要追加的slice 第二个是要追加的元素 最后返回一个 slice
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)
	s = append(s, "g")

	// 支持将一个 slice 复制到新的 sclice 中
	// 使用 copy 这个方法
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy", c)

	fmt.Println("len", len(s))
	fmt.Println("cap", cap(s))
	// 可以通过如下方式获取一个 slice 的片段 以至于生成一个新的 slice
	// 这里截取的长度是前闭后开 因为数组是以0开始计数的，这样做便于理解
	l := s[2:5]
	fmt.Println("sl1", l)

	l = s[:5]
	fmt.Println("sl2", l)

	l = s[2:]
	fmt.Println("sl3", l)

	// 以上介绍的是声明 slice 之后，采用 append 填充值
	// 接下来就是初始化的时候赋值。这里和数组的区别就是不需要指定长度
	// var a [3]string 数组的声明方式
	t := []string{"1", "2", "3"}
	fmt.Println("dclear", t)

	// 同时也支持二维 slice 一个大数组中有3个小数组 内部的长度可以变化，不像多维数组
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLine := i + 1
		twoD[i] = make([]int, innerLine)
		for j := 0; j < innerLine; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
}
