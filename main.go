package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache, no-store")
	fmt.Fprint(w, "<h1>Welcome to my Amazing Site !!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache, no-store")
	fmt.Fprintf(w, "<h1>Contact Page </h1><p>To get in touch, email me at <a href=\"mailto:ojha.saurabh10@gmail.com\">Saurabh Ojha</a></p>")
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/contact", contactHandler)
	fmt.Println("Starting the server on :3001...")
	http.ListenAndServe(":3001", nil)
}
