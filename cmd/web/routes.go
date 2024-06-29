package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	// Route pattern that ends with trailing slash -- like "/" or "/static/" --
	// it is known as a 'subtree path pattern', they act like "/**" or "/static/**"
	//
	// To prevent subtree path pattern, append `{$}` -- like "/{$}" or "/static/{$}"
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
	mux.HandleFunc("GET /snippet/create", app.snippetCreate)
	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)

	return commonHeaders(mux)
}
