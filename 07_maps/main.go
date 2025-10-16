package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int)

	m["one"] = 1
	m["two"] = 2
	m["three"] = 3

	fmt.Println("Map:", m)
	delete(m, "two")
	fmt.Println("Map:", m)

	clear(m)
	fmt.Println("Map after clear:", m)

	m2 := map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println("Map 2:", m2)

	val, ok := m2["two"]
	if ok {
		fmt.Println("Value for 'two':", val)
	} else {
		fmt.Println("'two' not found in map")
	}

	fmt.Println(maps.Equal(m, m2))
}
