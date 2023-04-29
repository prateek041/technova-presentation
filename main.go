package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "This is the Hello")
}

func GoodbyeHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "This is the")
}

func main() {
	http.HandleFunc("/hello", HelloHandler)
	http.HandleFunc("/goodbye", GoodbyeHandler)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("There is some error in starting the server", err)
	}
}
