package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("channels in golang - github.com")

	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}

	// fmt.Println(<-myCh)
	// myCh <- 5
	wg.Add(2)

	// Receive ONLY channel
	go func(ch <-chan int, wg *sync.WaitGroup) {
		close(myCh)
		val, isChannelOpen := <-myCh
		fmt.Println(isChannelOpen)
		fmt.Println(val)
		//fmt.Println(<-myCh)

		wg.Done()
	}(myCh, wg)

	// send ONLY channel
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myCh <- 0
		// myCh <- 5
		// myCh <- 6
		close(myCh)
		wg.Done()
	}(myCh, wg)
	wg.Wait()
}
