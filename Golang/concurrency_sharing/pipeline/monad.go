package main

import (
	"fmt"
)

func main() {
	intSteamGenerator(1, 2, 3, 5).multiplierStage(3).multiplierStage(5).printDatasink()
}

type intStream <-chan int

// Generators
func intSteamGenerator(integers ...int) intStream {
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
func (input intStream) multiplierStage(multiplier int) intStream {
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
func (input intStream) printDatasink() {
	for i := range input {
		fmt.Println(i)
	}
}
