package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/tasks", app.listTasks)
	mux.HandleFunc("/tasks/view", app.showTask)
	mux.HandleFunc("/tasks/create", app.createTask)

	return mux
}
