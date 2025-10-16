package main

import (
	"fmt"
	"sync"
)

func task(id int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("doing task", id)
}

func main() {

	var waitGP sync.WaitGroup

	for i := 0; i <= 10; i++ {
		waitGP.Add(1)
		go task(i, &waitGP)
	}
	waitGP.Wait()

}
