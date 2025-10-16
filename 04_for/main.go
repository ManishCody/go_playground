package main

import "fmt"

func main() {
	i := 1

	for i <= 5 {
		fmt.Println("Hello, World!")
		i = i + 1
	}

	for j := 1; j <= 5; j++ {
		fmt.Println("Hello, World!")
	}

	if i > 5 {
		fmt.Println("i is greater than 5")
	} else {
		fmt.Println("i is not greater than 5")
	}

}
