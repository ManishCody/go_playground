package main

import "fmt"

func printSlice[T any](item []T) {
	for _, it := range item {
		fmt.Println((it))
	}
}

func printIntStringSliceOnly[T int | string](item []T) {
	for _, it := range item {
		fmt.Println((it))
	}
}

type stack[T any] struct {
	ele []T
}

func main() {
	item := []int{1, 2, 4, 5, 6, 78}
	names := []string{"mandy", "sandy", "landy", "pandy"}

	printSlice(item)
	printIntStringSliceOnly(names)

	myStack := stack[int]{
		ele: []int{1, 123, 123, 213, 21, 312, 3, 12, 3},
	}

	fmt.Println(myStack)

}
