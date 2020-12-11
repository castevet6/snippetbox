package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
    // prevent catch all
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

    // init a slice containing two template files
    files := []string{
        "./ui/html/home.page.tmpl",
        "./ui/html/footer.partial.tmpl",
        "./ui/html/base.layout.tmpl",
    }
    
    // use parse home page template, return error or page
    // note variadic param files...
    ts, err := template.ParseFiles(files...)
    if err != nil {
        log.Println(err.Error())
        http.Error(w, "Internal Server Error", 500)
        return
    }

    // write tmpl content to response body with Execute(), no w.Write needed:wq
    // no dynamic data (last param of Execute)
    err = ts.Execute(w, nil)
    if err != nil {
        log.Println(err.Error())
        http.Error(w, "Internal Server Error", 500)
    }
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil || id < 1 {
        http.NotFound(w, r)
        return
    }

    fmt.Fprintf(w, "Display snippet: %d", id)
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        w.Header().Set("Allow", http.MethodPost)
        http.Error(w, "Method not allowed", 405)
        return
    }

    w.Write([]byte("Create new snippet..."))
}
