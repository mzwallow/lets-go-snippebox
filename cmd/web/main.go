package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	// Route pattern that ends with trailing slash -- like "/" or "/static/" --
	// it is known as a 'subtree path pattern', they act like "/**" or "/static/**"
	//
	// To prevent subtree path pattern, append `{$}` -- like "/{$}" or "/static/{$}"
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Print("starting server on :4000")

	// If using ":http" or named port, the `http.ListenAndServe()` function will try to lookup revelant
	// port number from `/etc/services`
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
