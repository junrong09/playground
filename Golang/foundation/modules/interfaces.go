package modules

import (
	"fmt"
)

func RunInterfaces() {
	var p Printable = &Point{1,2} // Interface in Go are implicit. (No need to declare implements etc...)
	//p = Vertex{1,2} // Error, as Vertex does not have method "Print". (Only *Vertex has)

	p.Print()
	//p.Reset() // Error, Does not exist

	// Essentally when a method is called on a interface, the method of the underlying type will be executed instead
	fmt.Printf("%v, %T\n", p, p) // val: &{1,2}; type: *modules.Point

	var nilP Printable
	fmt.Printf("%v, %T\n", nilP, nilP) // val: nil, type: nil
	//nilP.Print() // Run-time error when method called

	var p2 *Point
	nilP = p2
	fmt.Printf("%v, %T\n", nilP, nilP) // val: nil; type: *modules.Point
	nilP.Print() // Run-time error, if method does not check nil

	describe([]int{1,2}) //[1,2] []int
	assertion()
}

type Printable interface {
	Print()
	ToString() string
}

type Point struct {
	x,y int
}

func (v *Point) Print() {
	// Method should handle nil struct in Go
	if v == nil {
		return
	}
	fmt.Println(*v)
}

func (v *Point) ToString() string {
	return string(v.x) + " " + string(v.y)
}

func (v *Point) Reset() {}


// Usage of empty interface
// Func now can take in any type
func describe(i interface{}) { // Can replace interface{} with emptyInterface. Its the same.
	fmt.Printf("%v %T\n", i, i)
}

type emptyInterface interface {}

// Interface concrete value type assertion
func assertion() {
	var i Printable = &Point{1,2}
	p, ok := i.(*Point)
	fmt.Println(p, ok) // &{1 2} true

	var ii interface{} = "hello"
	//s := ii.(float64) // panic, if 2nd return val not captured
	s,_ := ii.(float64)
	fmt.Println(s) // 0 (float64 zero-val)

	// Use to check type
	do(ii)
}

func do(i interface{}) {
	switch v := i.(type) { // i.(type) can only be used in a switch
	case string:
		fmt.Printf("%v is a string", v)
	default:
		fmt.Printf("%v is a unknown type", v)
	}
}