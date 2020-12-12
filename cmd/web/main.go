package main

import (
    "flag"
    "log"
    "net/http"
    "os"
)

type application struct {
    errorLog *log.Logger
    infoLog *log.Logger
}

func main() {
    // Define flag addr for HTTP address
    addr := flag.String("addr", ":4000", "HTTP Network Address")
    flag.Parse()
    
    // create info logger
    infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

    // create error log
    errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

    // initialize new instance of application containing the depndencies
    app := &application{
        errorLog: errorLog,
        infoLog: infoLog,
    }

    mux := http.NewServeMux()
    mux.HandleFunc("/", app.home)
    mux.HandleFunc("/snippet", app.showSnippet)
    mux.HandleFunc("/snippet/create", app.createSnippet)

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
