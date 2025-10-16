package main

func sum(num ...int) int {
	total := 0

	for _, it := range num {
		total += it
	}
	return total
}

func main() {
	nums := []int{1, 2, 324, 4332, 32}
	var res int = sum(nums...)
	println("The sum is:", res)
}
