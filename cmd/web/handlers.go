package main

import (
    "errors"
    "fmt"
    "html/template"
    "net/http"
    "strconv"

    "github.com/castevet6/snippetbox/pkg/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
    // prevent catch all
    if r.URL.Path != "/" {
        app.notFound(w) // use helper method
        return
    }

    s, err := app.snippets.Latest()
    if err != nil {
        app.serverError(w, err)
        return
    }

    for _, snippet := range s {
        fmt.Fprintf(w, "%v\n", snippet)
    }

    /*
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
    }*/
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil || id < 1 {
        app.notFound(w) // use helper method
        return
    }

    // use SnippetModels object's Get method
    s, err := app.snippets.Get(id)
    if err != nil {
        if errors.Is(err, models.ErrNoRecord) {
            app.notFound(w)
        } else {
            app.serverError(w, err)
        }
        return
    }

    data := &templateData{Snippet: s}

    // init slice containing paths to show.page.tmpl file, plus the base layout and partial
    files := []string{
        "./ui/html/show.page.tmpl",
        "./ui/html/base.layout.tmpl",
        "./ui/html/footer.partial.tmpl",
    }

    // parse template files
    ts, err := template.ParseFiles(files...)
    if err != nil {
        app.serverError(w, err)
        return
    }
    
    // Execute template, pass snippets with parameter data
    err = ts.Execute(w, data)
    if err != nil {
        app.serverError(w, err)
    }
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        w.Header().Set("Allow", http.MethodPost)
        http.Error(w, "Method not allowed", 405)
        return
    }

    // mock data
    title := "O snail"
    content := "O snail\nClimb Mount Fuji\nBut slowly, slowly!\n\n- Kobayashi Issa"
    expires := "7"

    // pass data to SnippetModel.Insert() method, return ID
    id, err := app.snippets.Insert(title, content, expires)
    if err != nil {
        app.serverError(w, err)
        return
    }

    http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}
