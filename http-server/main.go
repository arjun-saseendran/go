package main

import (
	"fmt"
	"net/http"
)

func main() {
server := &http.Server{
	Addr: ":3000",
}
http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Ok")
})
server.ListenAndServe()
}

