package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

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

	logger.Info("starting server", "addr", *addr)

	// If using ":http" or named port, the `http.ListenAndServe()` function will try to lookup revelant
	// port number from `/etc/services`
	err := http.ListenAndServe(*addr, mux)
	logger.Error(err.Error())
	os.Exit(1)
}
