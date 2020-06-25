package main

import (
	"errors"
	"fmt"
)

// 一个带有 error 的方法 by conversation erro 最后一个返回值
func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}

	return arg + 3, nil
}

// 自定义错误
type argError struct {
	arg     int
	problem string
}

// 使用自定义的错误类型 实现 Error() 方法就是继承了 errors 这个接口
// 这里查了一下官方文档 Printf 和 Sprintf 只是在返回值的部分有区别
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.problem)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with arg"}
	}

	return arg + 3, nil
}

func main() {

	for _, v := range []int{7, 42} {

		// 这里可以同时处理两个条件
		if r, e := f1(v); e != nil {
			fmt.Println("f1 failed: ", e)
		} else {
			fmt.Println("f1 worked: ", r)
		}
	}

	for _, v := range []int{7, 42} {

		// 这里可以同时处理两个条件
		if r, e := f2(v); e != nil {
			fmt.Println("f2 failed: ", e)
		} else {
			fmt.Println("f2 worked: ", r)
		}
	}

	// 断言语法 e.(Type) 返回当前实例 和 是否存在
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.problem)
	}
}
