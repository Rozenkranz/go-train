package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (a *app) Home(w http.ResponseWriter, r *http.Request) {
	a.InfoLog.Printf("Home was visited by %s\n with %s method\n", r.RemoteAddr, r.Method)

	if r.URL.Path != "/" {
		a.notFound(w)
		return
	}

	var files = []string{
		"./ui/html/home_page.tmpl",
		"./ui/html/base_layout.tmpl",
		"./ui/html/footer_partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		a.serverError(w, err)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		a.serverError(w, err)
	}
}

func (a *app) Quests(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		a.notFound(w)
		return
	}

	fmt.Fprintf(w, "You wanna see quests with %d ID", id)

	a.InfoLog.Printf("Quests was visited by %s\n with %s method\n", r.RemoteAddr, r.Method)
}

func (a *app) AddQuest(w http.ResponseWriter, r *http.Request) {
	//Проверка на тип запроса
	if r.Method != http.MethodPost {
		a.InfoLog.Printf("AddQuest was visited by %s\n with %s method\n", r.RemoteAddr, r.Method)
		a.clientError(w, http.StatusBadRequest)
		return
	}
	w.Write([]byte("Creating a new quest"))
	a.InfoLog.Printf("AddQuest was visited by %s\n with %s method\n", r.RemoteAddr, r.Method)
}
