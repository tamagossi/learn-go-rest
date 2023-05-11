package basic_routing

import (
	"fmt"
	"math/rand"
	"net/http"
)

type CustomServeMux struct{}

func (mux *CustomServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		giveRandom(w, r)
		return
	}

	http.NotFound(w, r)
	return
}

func giveRandom(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Your random number is: %f", rand.Float64())
}

func CustomMux() {
	mux := &CustomServeMux{}
	http.ListenAndServe(":8000", mux)
}
