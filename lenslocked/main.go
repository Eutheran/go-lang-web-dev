package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
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

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Page not found", http.StatusNotFound)
}

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(notFoundHandler)
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}

// Notes:
// http.Handler - interface with the ServeHTTP method
// http.HandlerFunc - a function type that accepts same args as ServeHttp method,
// also implements http.Handler
