package main

import "fmt"

func main() {
	var nums [5]int

	fmt.Println("Initial array:", nums)
	nums[0] = 10
	nums[1] = 20
	nums[2] = 30
	nums[3] = 40
	nums[4] = 50

	fmt.Println("Updated array:", nums)

	// Accessing elements
	fmt.Println("First element:", nums[0])
	fmt.Println("Last element:", len(nums))

	//string array
	var fruits [3]string

	fmt.Println("Initial array:", fruits)
	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[2] = "Cherry"

	nums2 := [5]int{1, 2, 3, 4, 5}

	fmt.Println("Fruits array:", nums2)

	matrix := [2][2]int{{3, 4}, {1, 2}}
	fmt.Println("Matrix:", matrix)

}
