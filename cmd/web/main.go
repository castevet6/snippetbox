package main

import (
    "database/sql"
    "flag"
    "log"
    "net/http"
    "os"

    "github.com/castevet6/snippetbox/pkg/models/mysql"

    _ "github.com/go-sql-driver/mysql"
)

type application struct {
    errorLog *log.Logger
    infoLog  *log.Logger
    snippets *mysql.SnippetModel
}

func main() {
    // Define flag addr for HTTP address
    addr := flag.String("addr", ":4000", "HTTP Network Address")
    // define DSN for database
    dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")
    flag.Parse()
    
    // create info logger
    infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
    errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

    // open db conn pool via openDB wrapper function
    db, err := openDB(*dsn)
    if err != nil {
        errorLog.Fatal(err)
    }
    defer db.Close()

    // initialize new instance of application containing the depndencies
    app := &application{
        errorLog: errorLog,
        infoLog: infoLog,
        snippets: &mysql.SnippetModel{DB: db},
    }

    // initialize server struct
    srv := &http.Server {
        Addr: *addr,
        ErrorLog: errorLog,
        Handler: app.routes(), // return mux from routes.go
    }

    infoLog.Printf("Starting server on %s\n", *addr)
    err = srv.ListenAndServe()
    errorLog.Fatal(err)
}

// wrapper for sql.Open with given DSN
func openDB(dsn string) (*sql.DB, error) {
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }
    if err = db.Ping(); err != nil {
        return nil, err
    }
    return db, nil
}
