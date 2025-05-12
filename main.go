package main

import (
	"fmt"
	"net/http"
)

/* type login int
type welcome int

func (l login) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "on login page")
}

func (wl welcome) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "on welcome page")
} */

/* type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
} */

func myLogin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprint(w, "Using GET")
	case "POST":
		fmt.Fprintln(w, "Using POST")
	}

	fmt.Println(w, "on login page")
}

func myWelcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "My Request : %+v\n", r)
	fmt.Println(w, "on login page")
}

func main() {
	http.HandleFunc("/login", myLogin)
	http.HandleFunc("/welcome", myWelcome)
	// http.Handle("/login", http.HandlerFunc(myLogin))
	// http.Handle("/welcome", http.HandlerFunc(myWelcome))
	// var i login
	// var j welcome
	// http.Handle("/login", i)
	// http.Handle("/welcome", j)
	fmt.Println("Listening on port 8080...")
	http.ListenAndServe("localhost:8080", nil)
}
