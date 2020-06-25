package main

import "fmt"

// 这里是有关迭代循环的东西
func main() {
	// slice
	nums := []int{2, 3, 4}
	sum := 0

	for index, num := range nums {
		fmt.Println(index)
		sum += num
	}

	fmt.Println(sum)

	kvs := map[string]string{"a": "apple", "b": "banana"}

	// 遍历 map 的话。返回的两个值一个是 key  一个是 value
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// 如果返回译者，那么就是 key
	for k := range kvs {
		fmt.Println("key", k)
	}

	// 这里的情况和 array 或者 slice 有点像,这里是输出 unicode
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
