package main

import (
	"fmt"
	"net/http"
)

type login int
type welcome int

func (l login) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "on login page")
}

func (wl welcome) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "on welcome page")
}

/* type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
} */

func myLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `	
		<html>
		  <head>
		    Hi
		  </head>
	      <body>
		  <h1>
		   Please enter your username and password
		  </h1>
	      </body>
		<html>	
	`)
}

func myWelcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `	
		<html>
		  <head>
		    Hi
		  </head>
	      <body>
		  <h1>
		   Welcome
		  </h1>
	      </body>
		<html>	
	`)
}

func main() {
	// http.HandleFunc("/login", myLogin)
	// http.HandleFunc("/welcome", myWelcome)
	// http.Handle("/login", http.HandlerFunc(myLogin))
	// http.Handle("/welcome", http.HandlerFunc(myWelcome))
	var i login
	var j welcome
	http.Handle("/login", i)
	http.Handle("/welcome", j)
	fmt.Println("Listening on port 8080...")
	http.ListenAndServe("localhost:8080", nil)
}
