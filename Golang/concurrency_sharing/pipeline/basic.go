package main

import (
	"fmt"
)

func main() {
	intGenerator := IntSteamGenerate(1, 2, 3, 4)
	resultStream := multiplierStage(multiplierStage(intGenerator, 3), 5)
	PrintDatasink(resultStream)
}

// Generators
func IntSteamGenerate(integers ...int) <-chan int {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for _, i := range integers {
			intStream <- i
		}
	}()
	return intStream
}

// Stages
func multiplierStage(input <-chan int, multiplier int) <-chan int {
	multipliedStream := make(chan int)
	go func() {
		defer close(multipliedStream)
		for i := range input {
			multipliedStream <- i * multiplier
		}
	}()
	return multipliedStream
}

// Data Sinks
func PrintDatasink(input <-chan int) {
	for i := range input {
		fmt.Println(i)
	}
}
