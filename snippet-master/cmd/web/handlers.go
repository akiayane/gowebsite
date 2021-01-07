package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		//http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		//http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
		//http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		//http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display Snippet with ID %v", id)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		//http.Error(w, "Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create Snippet"))
}
