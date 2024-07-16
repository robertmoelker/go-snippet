package main

import (
	"html/template"
	"net/http"
	"path"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Write([]byte("Hello, World!"))
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	w.Write([]byte("Display a specific snippet..."))
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create a new snippet..."))
}
