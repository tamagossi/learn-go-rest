package middleware_rpc

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/justinas/alice"
)

func filterContentType(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Currently in the check content type middleware")

		if r.Header.Get("Content-type") != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write([]byte("415 - Unsupported media type. Please send JSON"))
			return
		}

		handler.ServeHTTP(w, r)
	})
}

func setServerTimeCookie(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
		cookie := http.Cookie{Name: "Server-Time(UTC)", Value: strconv.FormatInt(time.Now().Unix(), 10)}
		http.SetCookie(w, &cookie)
		log.Println("Currently in the set server time middleware")
	})
}

// func main() {
// 	mainLogicHandler := http.HandlerFunc(mainLogic)
// 	http.Handle("/city",
// 		filterContentType(setServerTimeCookie(mainLogicHandler)))

// 	http.ListenAndServe(":8000", nil)
// }

func MultipleMiddleware() {
	mainLogicHandler := http.HandlerFunc(logginMiddlewareLogic)
	chain := alice.New(filterContentType,
		setServerTimeCookie).Then(mainLogicHandler)
	http.Handle("/city", chain)
	http.ListenAndServe(":8000", nil)
}
