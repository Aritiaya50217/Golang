package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// แสดงข้อความหน้าเว็บ
		fmt.Fprintf(w, "Welcome to my website!")
	})
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// port
	http.ListenAndServe(":80", nil)

}
