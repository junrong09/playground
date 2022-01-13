package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//fmt.Println(runtime.NumCPU())
	//runtime.GOMAXPROCS(1)

	t1 := time.Now()
	randomGenerator(4).slowStage(time.Second * 2).dummyDatasink()
	fmt.Println("completes in ", time.Now().Sub(t1)) // completed in  8s

	t2 := time.Now()
	randomGenerator(4).fanout(4).concurrentSlowStage(time.Second * 2).fanin().dummyDatasink()
	fmt.Println("completes in ", time.Now().Sub(t2)) // completed in  2s / 4s / 6s / 8s
}

type interfaceStream <-chan interface{}

// Generators
func randomGenerator(n int) interfaceStream {
	output := make(chan interface{})
	go func() {
		defer close(output)
		for i := 0; i < n; i++ {
			output <- time.Now()
		}
	}()
	return output
}

// Stages
func (input interfaceStream) slowStage(timeToProcess time.Duration) interfaceStream {
	output := make(chan interface{})
	go func() {
		defer close(output)
		for i := range input {
			time.Sleep(timeToProcess)
			output <- i
		}
	}()
	return output
}

// Data Sinks
func (input interfaceStream) dummyDatasink() {
	for i := range input {
		elapse := time.Now().Sub(i.(time.Time))
		_ = elapse
		//fmt.Println(elapse)
	}
}

// utilities
func (input interfaceStream) fanout(fanOutCount int) interfaceStreams {
	outputs := make([]interfaceStream, fanOutCount)
	for i := 0; i < fanOutCount; i++ {
		output := make(chan interface{})
		outputs[i] = output
		go func() {
			defer close(output)
			for i := range input {
				output <- i
			}
		}()
	}
	return outputs
}

type interfaceStreams []interfaceStream

// Stages
func (inputs interfaceStreams) concurrentSlowStage(timeToProcess time.Duration) interfaceStreams {
	outputs := make([]interfaceStream, len(inputs))
	for i, input := range inputs {
		outputs[i] = input.slowStage(timeToProcess)
	}
	return outputs
}

// utilities
func (inputs interfaceStreams) fanin() interfaceStream {
	var wg sync.WaitGroup
	multiplexedOutput := make(chan interface{})

	multiplex := func(input <-chan interface{}) {
		defer wg.Done()
		for i := range input {
			multiplexedOutput <- i
		}
	}

	wg.Add(len(inputs))
	for _, input := range inputs {
		go multiplex(input)
	}

	go func() {
		wg.Wait()
		close(multiplexedOutput)
	}()

	return multiplexedOutput
}
