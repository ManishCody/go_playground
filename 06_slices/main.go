package main

import "fmt"

func main() {

	var nums []int

	fmt.Println("Initial slice:", len(nums))

	var nums2 = make([]int, 0, 10)
	fmt.Println(cap(nums2))
	fmt.Println("Initial slice:", nums2, len(nums2))
	nums2 = append(nums2, 1)
	nums2 = append(nums2, 2, 3, 4, 5, 6, 3)
	fmt.Println(cap(nums2))
	fmt.Println("Initial slice:", nums2, len(nums2))

	num3 := []int{}
	num3 = append(num3, 1, 2, 3, 4, 5)
	fmt.Println("Slice with values:", num3)

	var slice1 = make([]int, 0, 5)
	var slice2 = make([]int, len(slice1))

	slice1 = append(slice1, 1, 2, 3, 4, 5)

	copy(slice2, slice1)

	fmt.Println("Slice1:", slice1)
	fmt.Println("Slice2:", slice2)

	// Slicing an slice
	fmt.Println("Slice1[1:3]:", slice1[1:3])
	fmt.Println("Slice1[:3]:", slice1[:3])
	fmt.Println("Slice1[2:]:", slice1[2:])

	// Iterating over a slice
	for i, v := range slice1 {
		fmt.Printf("Index %d has value %d\n", i, v)
	}

}
