package modules

import (
	"fmt"
	"math"
)

func RunMethods() {
	v1 := Vertex{1,2}
	fmt.Println(v1.Abs())
	fmt.Println(Abs(v1))

	v1.scale(10)	// Even if the func receiver takes a ptr, no need to get addr. (e.g. (&v1).scale(..))
	(&v1).scale(10) // same as prev.

	v1.scale2(10)
	(&v1).scale2(10) //Ptr can pass into func receiver that takes in value
	fmt.Println(v1)
}

type Vertex struct {
	x, y float64
}

// No classes in Go. But can create method for struct using func
// This is called func with "receiver"
func (v Vertex)Abs() float64 {
	return math.Sqrt(v.x * v.x + v.y * v.y)
}

// Same as prev
func Abs(v Vertex) float64{
	return math.Sqrt(v.x * v.x + v.y * v.y)
}

// Ptr receiver is required to edit the instance
// Without it the struct will just be a new copy and the scale will not persist
// Ptr receiver more efficient, because it does not copy the attributes in the struct
func (v *Vertex) scale(s float64) {
	v.x *= s
	v.y *= s
}

// Bad method to edit vertex
func (v Vertex) scale2(s float64) {
	v.x *= s
	v.y *= s
}