package basics

import (
	"fmt"
	"rsc.io/quote"
)

func RunVariables() {
	var a int
	fmt.Println(quote.Go())
	a,b := multipleReturn() // b is not declared, so can use :=
	f := namedReturn(10)
	fmt.Println(f,a,b)
	typeInfer()
}

// Outside function, all statement must begin with keyword (e.g. func, const, var)

// Function return
func multipleReturn() (int, int) {
	var c = 20
	var a, b int = 1, c
	return a, b
}

func namedReturn(a int) (b int) {
	b = a ^ 2
	return
}

// Variables declaration and assignment
func declaration() {
	const A = 10 // const cannot be declared using :=
	var a, b int // Default int (0)
	var c,d = 10,20
	e,f := 30, 40 // short declaration with assignment (without var)
	var (
		g bool = false
	)
	fmt.Println(a,b,c,d,e,f,g)
}

func casting() float32 {
	return float32(10)
}

func typeInfer() {
	// Default type given a constant
	i := 42           // int
	f := 3.142        // float64
	g := 0.867 + 0.5i // complex128
	g2 := g // Infer from type of right variable
	fmt.Printf("Type: %T %T %T %T", i,f,g,g2)
}

func constants() {
	const A = 10
	const (
		big = 1 << 100 // Precise now, overflow when given context (type)
	)

	//namedReturn(big) // Constant will overflow when cast to int
}