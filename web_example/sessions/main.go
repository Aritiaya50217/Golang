package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func secret(w http.ResponseWriter, r *http.Request) {
	sessions, _ := store.Get(r, "cookie-name")
	// check if user is authenticated
	if auth, ok := sessions.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	// Print secret message
	fmt.Fprintln(w, "The cake is a lie !!!")
}

func login(w http.ResponseWriter, r *http.Request) {
	sessions, _ := store.Get(r, "cookie-name")
	// Authentication goes here
	// ...
	// set user as authenticated
	sessions.Values["authenticated"] = true
	sessions.Save(r, w)
}

func logout(w http.ResponseWriter, r *http.Request) {
	sessions, _ := store.Get(r, "cookie-name")
	// Revoke users authentication
	sessions.Values["authenticated"] = false
	sessions.Save(r, w)
}

func main() {
	http.HandleFunc("/secret", secret)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)

	// port
	http.ListenAndServe(":8080", nil)
}
