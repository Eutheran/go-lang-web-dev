package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Contact Page</h><p>To get in touch, email me at <a href=\"mailto:ocrashandburno@gmail.com\">ocrashandburno@gmail.com.</a>")
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		http.Error(w, "Page Not Found", http.StatusNotFound)
	}
}

// type Router struct{}

// func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	default:
// 		http.Error(w, "Page Not Found", http.StatusNotFound)
// 	}
// }

func main() {
	fmt.Println("Starting the server on :3000...")
	// we are not passing the pathHandler into a function
	// this http.HandlerFunc is converting the handler type
	http.ListenAndServe(":3000", http.HandlerFunc(pathHandler))
}

// Notes:
// http.Handler - interface with the ServeHTTP method
// http.HandlerFunc - a function type that accepts same args as ServeHttp method,
// also implements http.Handler
