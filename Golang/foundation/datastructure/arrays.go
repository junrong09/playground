package datastructure

import "fmt"

func RunArrays() {
	array()
	slice()
	sliceManipulation()
	makeSlice()
	appendSlice()
	iterate()
}

/**
Array
*/
func array() {
	var A [2]int // Array is size-fixed, must declare upfront
	A[0] = 10
	A[1] = 20
	change2(A)
	change3(&A)

	B := [3]int{1,2,3} // Without len in [], a slice will be created instead of a array
	C := [3]int{} //Must have {}
	D := B // Create copy of array B
	D[0] *= 2 // Does not affect B
	E := [...]int{1,2,3,4} //Let compiler count the len of array
	fmt.Println(A, B, C, D, E)
}

func change1(A []int) { // Function that takes in a slice
	A[0] = 100000
}

func change2(A [2]int) { // Function that takes in a array of size-2
	A[0] = 100000 // Since A is copied, original array not updated
}

func change3(A *[2]int) { //ptr updates the A[0] of array passed in
	A[0] = 100000
}

/**
Slice
 */
func slice() {
	nums := [6]int{1,2,3,4,5,6}

	var s1 []int = nums[1:4] //inclusive low, exclusive high
	s2 := nums[:] // low and high are optional, default to 0 and size respectively
	s2[1] = 10
	fmt.Println(s1, s2) //S1 and S2 are just refer/view to nums

	s3 := []int{1,2,3} // slice creation based on anonymous array
	s4 := []struct{ // Anonymous struct, initialized as slice with 2 elements
		i int
		b bool
	}{
		{2, true},
		{3, false},
	}

	var s5 []int // no underlaying array (== nil) is true
	fmt.Println(s3, s4, s5 == nil)
}

func sliceManipulation() {
	s := []int{1,2,3,4,5,6} // slice with underlaying array of {1,2,3,4,5,6}

	// IMPORTANT: Point of reference is always the start of the slice
	// len => length of slice
	// cap => size of underlaying array the slice is based on (counted from first element in slice)

	s = s[:0] // len 0 slice
	fmt.Println(len(s), cap(s)) //0,6

	s = s[:3] // len 3 slice
	fmt.Println(len(s), cap(s)) //3, 6

	//s = s[:7] // Error, exceeded capacity

	s = s[2:] // Drop first 2
	fmt.Println(len(s), cap(s)) //1, 4 (cap counts from start of slice to end)

	s = s[:cap(s)] // Extend to the end of cap
	fmt.Println(s)

	s2 := s // Create slice based on existing slice
	s2[0] = 1000
	fmt.Println(s, s2) // Same slice low and high, so same value
}

func makeSlice() {
	s := make([]int, 5) //set len
	// similar to that of, s := []byte{0, 0, 0, 0, 0}
	fmt.Println(len(s), cap(s))

	s = make([]int, 0, 5) //set len 0, cap 5
	fmt.Println(len(s), cap(s))
}

func appendSlice() {
	var s []int
	s = append(s, 0)
	fmt.Println(len(s), cap(s), s) //1, 1

	s = append(s, 1)
	fmt.Println(len(s), cap(s), s) //2, 2

	s = append(s, 2)
	fmt.Println(len(s), cap(s), s) //3, 4

	s = append(s, 3,4)
	fmt.Println(len(s), cap(s), s) //5, 8 (Array growth 2* of previous)
}

/**
Range/iterating
 */
func iterate() {
	var A []int = []int{1,2,3}
	for i,v := range A {
		fmt.Println(i, v)
	}
	for _,v := range A { // Ignore val of multiple return
		fmt.Println(v)
	}

	// Example of creating a 2D-slice
	B := make([][]int, 10)
	for i,_ := range B {
		B[i] = make([]int, 10)
	}
}
