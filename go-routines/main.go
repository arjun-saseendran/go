package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}
var wg sync.WaitGroup // pointers
var mut sync.Mutex    // pointers

func main() {
	// go	greet("Hello")
	//
	//	greet("World")
	websiteList := []string{
		"https://rust.dev",
		"https://go.dev",
		"https://node.dev",
	}
	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signals)
}

// func greet(text string) {
// 	for i := 0; i <= 5; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(text)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS endpoint")

	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}
