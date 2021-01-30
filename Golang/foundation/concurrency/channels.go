package concurrency

import (
	"fmt"
	"math/rand"
	"time"
)

func RunRoutines() {
	ChannelExample()
	ChannelRangeExample()
	ChannelSelectExample()
}

func ChannelExample() {
	s := []int{1,2,3,4,5}
	ch := make(chan int, 2) // Construct channel with 2 buffer spaces (i.e. cap(ch) == 2)
	// Divide and conquer

	go sum(s[:len(s)/2], ch)
	go sum(s[len(s)/2:], ch)

	// len(ch) == 2, once the computation is done
	x, y := <-ch, <-ch // Channel used to syn data
	fmt.Printf("%d + %d = %d (buffer used = %d/%d)\n", x, y, x+y, len(ch), cap(ch))
}

func sum(A []int, ch chan int) {
	acc := 0
	for _,v := range A {
		acc += v
	}
	ch <- acc
}

func ChannelRangeExample() {
	ch := make(chan int)
	go generateRandom(ch)
	for v := range ch { // loop until channel close
		fmt.Println("Print random: ", v)
	}
	v, ok := <- ch
	fmt.Println(v, ok)
}

func generateRandom(ch chan int) {
	ch <- rand.Int()
	ch <- rand.Int()
	ch <- rand.Int()
	close(ch) // Closing before sending will cause panic
}

func ChannelSelectExample() {
	ch, quit := make(chan int), make(chan int)
	go populate(ch, quit)
	for {
		select { // Note: not switch case
		case v := <- ch:
			fmt.Println(v)
		case <- quit:
			fmt.Println("QUITING")
			return
		default:
			fmt.Println("...")
			time.Sleep(3 * time.Second)
		}
	}
}

func populate(ch, quit chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	quit <- 1
}