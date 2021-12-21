package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Middleware คือ ซอฟต์แวร์กลางที่เชื่อม Client และ Server เข้าด้วยกัน
type Middleware func(http.HandlerFunc) http.HandlerFunc

// logging logs all requests with its path and the time it tppk to process
func Logging() Middleware {
	// create a new middleware
	return func(f http.HandlerFunc) http.HandlerFunc {
		// define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {
			// do middleware things
			strat := time.Now()
			defer func() { log.Println(r.URL.Path, time.Since(strat)) }()

			// call the next middleware/handler in chain
			f(w, r)
		}
	}
}

// method ensures that url can only be requested with a specific method , else returns a 400 Bad Request
func Method(m string) Middleware {
	// create a new middleware
	return func(f http.HandlerFunc) http.HandlerFunc {
		// define the the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {
			// do middleware things
			if r.Method != m {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
			// call the next middleware/handler in chain
			f(w, r)
		}
	}
}

// chain applies middlewares to a http.HandlerFunc
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "helloworld")
}

func main() {
	http.HandleFunc("/", Chain(Hello, Method("GET"), Logging()))
	http.ListenAndServe(":8080", nil)
}
