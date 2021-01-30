package basics

import "fmt"

func RunPointers() {
	ptr()
	object()
	objectCopy()
}

func ptr() {
	var (
		p *int // Create ptr variable
		i = 10
	)
	p = &i // generate pointer
	*p = 20 // access/modify ptr variable with *
	//p++ //Error, no pointer arithmetic like C
	fmt.Println(i) //20
	fmt.Println(*createPtr()) // ptr not garbage collected
}

func createPtr() *int {
	ptr := new(int) // Shorthand to create *int without declaring a var
	*ptr = 100
	return ptr
}

type ( //Create new type (e.g. struct, interface)
	Vertex struct {
		x int
		y int
	}
)

type AnotherFloat float64

func object() {
	var (
		v1 = Vertex{1,2}  // construct object
		v2 = Vertex{x: 1} // y: 0
		v3 = Vertex{}     // 0, 0
		p1 = &Vertex{1,2} // create *Vertex
	)
	var p2 *Vertex
	p2 = &v1

	(*p2).x = 2 // Access attribute
	p2.x = 2 // Shorthand (with explicit dereference)

	fmt.Println(v1,v2,v3, p1, p2)
}

func objectCopy() {
	v1 := Vertex{1,2}
	v2 := v1 // Shallow-copy

	if v1 == v2 {
		fmt.Println("Equal") // True, even when they are diff instances. Because checks is based on attribute
	}

	v2.x *= 2
	if v1 != v2 {
		fmt.Println("Not Equal") // Not equal because v1.x != v2.x
	}
}