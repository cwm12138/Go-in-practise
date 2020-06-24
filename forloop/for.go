package main

import "fmt"

func main() {

	i := 1
	// while
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// for
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// while(true)
	for {
		fmt.Println("loop one")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}

		fmt.Println(n)
	}

	// range channels or others
}
