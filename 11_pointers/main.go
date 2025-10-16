package main

func updateVar(val *int) {
	*val = 1000
	println("Value inside function:", *val)
}

func main() {

	var num int = 100
	println("Before function call:", num)
	updateVar(&num)
	println("After function call:", num)

}
