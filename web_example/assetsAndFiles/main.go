package main

import "net/http"

func main() {
	// เข้าถึง directory
	fs := http.FileServer(http.Dir("assets/"))
	// url
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// port
	http.ListenAndServe(":8080", nil)
}
