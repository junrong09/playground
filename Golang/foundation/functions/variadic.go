package functions

import "fmt"

func RunVariadic() {
	sum(1,2,3)
	nums := []int{1,2,3,4}
	sum(nums...) // Like javascript spread operator
}

func sum(nums ...int) { // Variadic
	fmt.Printf("%T", nums) // It is a slice, e.g. []int
}