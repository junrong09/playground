package main

import (
	"fmt"
	"foundation/basics"
	"foundation/packages"

	"example.com/randomModule/random"
	"example.com/randomModule/random/innerRandom"
	"testing"
)

func main() {
	//basics.RunVariables()
	basics.RunPointers()
	//basics.RunControlFlow()

	//datastructure.RunArrays()
	//datastructure.RunMaps()

	//modules.RunMethods()
	//modules.RunInterfaces()
	//modules.RunErrors()

	//concurrency.RunRoutines()
	//concurrency.RunMutex()

	//functions.RunHigherOrderFunctions()
	//functions.RunVariadic()

	packages.RunPackage()
	fmt.Println(packages.RandomNum)

	v := packages.Vertex{}
	fmt.Println(v)
	fmt.Println(TT2(10) > 2)

	i := random.RandomInt()
	j := innerRandom.RandomMore()
	fmt.Println(i, j)
}

type TT1 float64
type TT2 TT1