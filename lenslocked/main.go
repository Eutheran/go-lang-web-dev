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
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h><p>To get in touch, email me at <a href=\"mailto:ocrashandburno@gmail.com\">ocrashandburno@gmail.com.</a>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	fmt.Fprint(w, `<h1>FAQ Page</h1>
	<ul>
	<li>
	<b>Is there a free version?</b>
	A: Yes! We offer a free trial for 30 days on any paid plans.</li>
	<li>
	<b>What are you support hours?</b>
	We have support staff answering emails and will try the best we can.</li>
	<li>
	<b>How do I contact support?</b>
	A: Email us --- just kidding, please don't email us.</li>
	</ul>
	`)
}

// func pathHandler(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	default:
// 		http.Error(w, "Page Not Found", http.StatusNotFound)
// 	}
// }

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		http.Error(w, "Page Not Found", http.StatusNotFound)
	}
}

func main() {
	var router Router
	fmt.Println("Starting the server on :3000...")
	// we are not passing the pathHandler into a function
	// this http.HandlerFunc is converting the handler type
	http.ListenAndServe(":3000", router)
}

// Notes:
// http.Handler - interface with the ServeHTTP method
// http.HandlerFunc - a function type that accepts same args as ServeHttp method,
// also implements http.Handler
