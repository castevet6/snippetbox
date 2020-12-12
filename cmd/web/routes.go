package main

import "net/http"

func (app *application) routes() *http.ServeMux {
    mux := http.NewServeMux()
    mux.HandleFunc("/", app.home)
    mux.HandleFunc("/snippet", app.showSnippet)
    mux.HandleFunc("/snippet/create", app.createSnippet)

    // create file server for static files
    fileServer := http.FileServer(http.Dir("./ui/static/"))

    // register fileserver for routes starting with /static
    mux.Handle("/static/", http.StripPrefix("/static", fileServer))

    return mux
}
