package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"

	"github.com/google/uuid"
	"github.com/robertmoelker/lets-go/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	tasks, err := app.tasks.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/view.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := &templateData{Tasks: tasks}

	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) showTask(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.URL.Query().Get("id"))
	if err != nil {
		app.notFound(w)
		return
	}

	task, err := app.tasks.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrRecordNotFound) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
	}

	fmt.Fprintf(w, "%+v", task)

	// w.Write([]byte("Display a specific Task..."))
}

func (app *application) createTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create a new Task..."))
}

func (app *application) listTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := app.tasks.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	for _, task := range tasks {
		fmt.Fprintf(w, "%+v\n", task)
	}
	// fmt.Fprintf(w, "%+v", tasks)

	// app.infoLog.Println(tasks)
	w.Write([]byte("Display the list of Tasks..."))
}
