package packages

import "fmt"

func RunPackage() {
	fmt.Println(randomNum) // Able to variable not in same go file, but in same package
	helperFn(10)

	var i int
	p := &i
	var pp = &p
	fmt.Printf("%T", pp)

	vp := new(Vertex)
	vp.x = 10
	xp := &vp.x
	*xp = 20
	fmt.Println(*vp)
}
