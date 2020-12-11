package main

import (
    "log"
    "net/http"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", home)
    mux.HandleFunc("/snippet", showSnippet)
    mux.HandleFunc("/snippet/create", createSnippet)

    // create file server for static files
    fileServer := http.FileServer(http.Dir("./ui/static/"))

    // register fileserver for routes starting with /static
    mux.Handle("/static/", http.StripPrefix("/static", fileServer))

    log.Println("Starting server on :4000")
    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)
}
