package main

import (
	"fmt"
	"sync"
)

func main() {
	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChanelOpen := <-myCh
		fmt.Println(isChanelOpen)
		fmt.Println(val)
		wg.Done()
	}(myCh, wg)
	go func(ch chan<- int, wg *sync.WaitGroup) {

		myCh <- 6
		close(myCh)
		wg.Done()
	}(myCh, wg)
	wg.Wait()

}
