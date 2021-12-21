package main

import (
	"fmt"
	"log"
	"net/http"
)

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "foo")
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "bar")
}

func say(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Gopher !!!")
}

func main() {
	/* This example will show how to create basic logging middleware in Go.
	A middleware simply takes a http.HandlerFunc as one of its parameters, wraps it and returns a new http.HandlerFunc for the server to call.
	*/
	http.HandleFunc("/foo", logging(foo))
	http.HandleFunc("/bar", logging(bar))
	http.HandleFunc("/say", logging(say))

	http.ListenAndServe(":8080", nil)
}
