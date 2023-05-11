package middleware_rpc

import (
	"fmt"
	"net/http"
)

func middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executing middleware before request phase!")
		handler.ServeHTTP(w, r)
	})
}

func customMiddlewareLogic(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Executing main handler...")
	w.Write([]byte("OK"))
}

func CustomMiddleware() {
	mainLogicHandler := http.HandlerFunc(customMiddlewareLogic)
	http.Handle("/", middleware(mainLogicHandler))
	http.ListenAndServe(":8000", nil)
}
