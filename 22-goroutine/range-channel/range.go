package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "1"
	queue <- "2"
	close(queue)

	for element := range queue {
		fmt.Println(element)
	}
}
