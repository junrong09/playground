package packages

import "math/rand"

var randomNum = rand.Int() // Accessible package wide only (see goFile1.go)
var RandomNum = rand.Int() // Accessible publicly via importing "packages"
type Vertex struct {
	x,y int
}

func helperFn(a int) int{
	return a * 2
}
