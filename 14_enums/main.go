package main

import "fmt"

type myType string

type OrderStatus int

const (
	Received OrderStatus = iota
	Confirmed
	Shipped
	Delivered
	Canceled
)

func changeStatus(status OrderStatus) {
	fmt.Println("Status inside method:", status)
}
func main() {

	changeStatus(Received)

}
