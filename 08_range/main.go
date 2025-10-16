package main

import "fmt"

func main() {
	num := []int{1, 2, 3, 4, 5, 6, 7, 8}
	sum := 0
	for i, n := range num {
		sum += n
		if i%2 == 0 {
			n *= 2
		} else {
			n *= 3
		}
		fmt.Println(n)
	}
	fmt.Println("Sum:", sum)

	map1 := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	for key, val := range map1 {
		fmt.Printf("Key: %s, Value: %d\n", key, val)
	}
}
