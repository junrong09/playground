package basics

import "fmt"

func RunControlFlow() {
	loops()
	conditions()
	switchs()
	deferStatement()
}

func loops() {
	for i := 1; i < 100; i++ { // No () around condition , but must have {}
		i++
	}

	j := 10
	for ;j < 20; { //init and increment are optional
		j ++
	}

	for j < 30 { //No "while" loop in Go, use "for" loop
		j++
	}
}

func conditions() {
	a := 10
	if a < 10 {

	}

	if b := 20; a < 20 { //Allow pre-condition execution, but b is only available in conditonal scope
		a += b
	} else if a > 100 {
		a -= b
	} else {
		a = b
	}
}

func switchs() {
	// Switch with optional pre-execution and expression
	switch a := 1; a {
	case 1:
		b := 2
		a += b
	default:
		a = 0
	}

	b := 10
	switch { // Switch without expression
	case b == 10:
		fmt.Println("Equals 10")
	case b > 10:
		fmt.Println("More than 10")
	default:
		fmt.Println("Less than 10")
	}
}

func deferStatement() {
	a := 10
	defer fmt.Println(a) //Parameters is evaluated immediately (i.e. 10), but executed when return is called
	defer fmt.Println(a + 1) // Last-in first-out defer statement execution (like a stack)
	a = 20
	fmt.Println(a)
	return // prints 20, 11, then 10.
}