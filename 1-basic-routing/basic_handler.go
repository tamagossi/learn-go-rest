package main

import (
	"io"
	"log"
	"net/http"
)

func Greeting(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!")
}

func main() {
	http.HandleFunc("/", Greeting)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
