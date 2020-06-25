package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	fmt.Println("iptr", iptr)

	*iptr = 0
}

func main() {

	i := 1
	fmt.Println("i:", i)

	zeroval(i)
	fmt.Println("i:", i)

	// & get i memory address
	zeroptr(&i)
	fmt.Println("i:", i)

	fmt.Println("i:", &i)

}
