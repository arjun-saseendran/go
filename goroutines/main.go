package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)

	go greet("Nice to meet you!", done)

	go greet("How are you?", done)

	go slowGreet("How... are... you...?", done)

	go greet("I hope you are liking the course!", done)

	for range done {

	}

}

func greet(phrase string, done chan bool) {
	fmt.Println("Hello!", phrase)
	done <- true
}

func slowGreet(phrase string, done chan bool) {
	time.Sleep(2 * time.Second)
	fmt.Println("Hello!", phrase)
	done <- true
	close(done)
}
