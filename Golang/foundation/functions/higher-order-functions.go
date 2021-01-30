package functions

import (
	"fmt"
	"math"
)

func RunHigherOrderFunctions() {
	pytha := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(pytha(5,12)) // call directly
	fmt.Println(compute(pytha)) // Pass to compute and it will call it

	addr := closure()
	addr(100)
	fmt.Println(addr(100)) // 200

	fibFn := fib()
	fmt.Println("Start of fib:")
	for i := 0; i < 10; i++ {
		fmt.Println(fibFn())
	}
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3,4)
}

func closure() func(int) int {
	sum := 0
	return func(a int) int {
		sum += a
		return sum
	}
}

func fib() func() int {
	a, b := 0,1
	return func() int {
		temp := a
		a = b
		b = temp + a
		return temp
	}
}