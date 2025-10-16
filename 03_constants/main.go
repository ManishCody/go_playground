package constants

func main() {
	const pi = 3.14
	const pi2 float32 = 3.14
	const pi3 string = "3.14"

	const (
		port = 8080
		host = "localhost"
	)

	// pi = 3.14159 // error: cannot assign to pi (constant variable)

	println("Value of pi:", pi)
	println("Value of pi2:", pi2)
	println("Value of pi3:", pi3)
	println("Server running at", host, "on port", port)
}
