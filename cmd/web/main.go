package main

import (
    "flag"
    "log"
    "net/http"
    "os"
)

func main() {
    // Define flag addr for HTTP address
    addr := flag.String("addr", ":4000", "HTTP Network Address")
    flag.Parse()
    
    // create info logger
    infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

    // create error log
    errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
    mux := http.NewServeMux()
    mux.HandleFunc("/", home)
    mux.HandleFunc("/snippet", showSnippet)
    mux.HandleFunc("/snippet/create", createSnippet)

    // create file server for static files
    fileServer := http.FileServer(http.Dir("./ui/static/"))

    // register fileserver for routes starting with /static
    mux.Handle("/static/", http.StripPrefix("/static", fileServer))

    // initialize server struct
    srv := &http.Server {
        Addr: *addr,
        ErrorLog: errorLog,
        Handler: mux,
    }

    infoLog.Printf("Starting server on %s\n", *addr)
    err := srv.ListenAndServe()
    errorLog.Fatal(err)
}
