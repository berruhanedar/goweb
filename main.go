package main

import (
	"fmt"
	"net/http"
)

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
	http.HandleFunc("/login", myLogin)
	http.HandleFunc("/welcome", myWelcome)
	fmt.Println("Listening on port 8080...")
	http.ListenAndServe("localhost:8080", nil)
}
