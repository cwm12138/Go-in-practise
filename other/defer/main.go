package main

import "fmt"

func second() {
	fmt.Println("two")
}

func first() {
	fmt.Println("one")
}

func main() {
	defer second()
	first()
}
