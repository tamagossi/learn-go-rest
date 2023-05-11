package basic_routing

import (
	"io"
	"log"
	"net/http"
)

func Greeting(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!")
}

func BasicHandler() {
	http.HandleFunc("/", Greeting)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
