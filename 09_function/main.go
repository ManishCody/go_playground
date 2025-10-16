package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func getLanguage() (string, string, int) {
	return "Go", "Python", 2023
}

func fn(a, b int) int {
	return a + b
}

func process(fn func(int, int) int) int {
	return fn(1, 2)
}

func main() {
	result := add(5, 10)
	println("The sum is:", result)
	language, _, year := getLanguage()
	println("Language:", language, "Year:", year)

	fmt.Println(process(fn))
}
