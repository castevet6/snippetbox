package main

import (
    "fmt"
    "html/template"
    "net/http"
    "strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
    // prevent catch all
    if r.URL.Path != "/" {
        app.notFound(w) // use helper method
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
        app.serverError(w, err) // use helper method
    }

    // write tmpl content to response body with Execute(), no w.Write needed:wq
    // no dynamic data (last param of Execute)
    err = ts.Execute(w, nil)
    if err != nil {
        app.serverError(w, err) // use helper method
    }
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil || id < 1 {
        app.notFound(w) // use helper method
        return
    }

    fmt.Fprintf(w, "Display snippet: %d", id)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        w.Header().Set("Allow", http.MethodPost)
        http.Error(w, "Method not allowed", 405)
        return
    }

    w.Write([]byte("Create new snippet..."))
}
