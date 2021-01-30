package datastructure

import (
	"fmt"
	"strings"
)

func RunMaps() {
	maps()
}

type Coordinate struct {
	x, y float32
}

func maps() {
	m1 := make(map[string]Coordinate)
	m1["lab1"] = Coordinate{1.23, 492.12}
	fmt.Println(m1["lab1"].x)

	var m2 = map[string]Coordinate{
		"lab1": Coordinate{1,2},
		"lab2": {2,3}, // Possible to omit the type (e.g. Coordinate)
	}
	fmt.Println(m2)

	m2["lab3"] = Coordinate{4,5} // insert/update
	delete(m2, "lab3")           // remove
	coordi, exist := m2["lab3"]  // Does not exist, return (zero-value {0,0}, bool: false)
	fmt.Println(coordi, exist)

	count := wordCount("test the word count")
	fmt.Println(count)
}

func wordCount(s string) map[string]int {
	var count = make(map[string]int)
	for _,v := range strings.Fields(s) {
		count[v] = count[v] + 1
	}
	return count
}