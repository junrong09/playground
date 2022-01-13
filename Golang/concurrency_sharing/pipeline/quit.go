package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	quitChan := make(chan interface{})

	// to simulate termination after some time
	go func() {
		time.Sleep(time.Second * 3)
		close(quitChan)
	}()

	t1 := time.Now()
	infSteamGenerator(quitChan, time.Millisecond*500).multiplierStage(3).printDatasink()
	fmt.Println("completes in ", time.Now().Sub(t1)) // completes in 3s but not always the case
}

type intStreamWithQuit struct {
	intStream <-chan int
	quitChan  <-chan interface{}
}

// Generators
func infSteamGenerator(quit <-chan interface{}, frequency time.Duration) intStreamWithQuit {
	output := make(chan int)
	go func() {
		defer close(output)
		for {
			time.Sleep(frequency)
			select {
			case output <- rand.Int():
			case <-quit:
				return
			}
		}
	}()
	return intStreamWithQuit{
		intStream: output,
		quitChan:  quit,
	}
}

// Stages
func (input intStreamWithQuit) multiplierStage(multiplier int) intStreamWithQuit {
	multipliedStream := make(chan int)
	go func() {
		defer close(multipliedStream)
		for {
			select {
			case i := <-input.intStream:
				multipliedStream <- i * multiplier
			case <-input.quitChan:
				return
			}
		}
	}()
	return intStreamWithQuit{
		intStream: multipliedStream,
		quitChan:  input.quitChan,
	}
}

// Data Sinks
func (input intStreamWithQuit) printDatasink() {
	for i := range input.intStream {
		fmt.Println(i)
	}
}
