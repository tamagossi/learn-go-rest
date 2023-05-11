package basic_routing

import (
	"fmt"
	"math/rand"
	"net/http"
)

func MultipleHandle() {
	mux := http.NewServeMux()

	mux.HandleFunc("/randomFloat", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, rand.Float64())
	})

	mux.HandleFunc("/randomInt", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, rand.Intn(100))
	})

	http.ListenAndServe(":8000", mux)
}
