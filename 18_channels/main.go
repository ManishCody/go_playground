package main

import (
	"fmt"
	"time"
)

func processNum(numChan chan int) {
	fmt.Println("processing number", <-numChan)
}

func main() {

	numChan := make(chan int)

	go processNum(numChan)

	numChan <- 5
	time.Sleep(time.Second * 2)
	// messageChan := make(chan string)

	// messageChan <- "Mandy"

	// msg := <-messageChan

	// fmt.Println(msg)
}
