package main

import (
	"net/http"
)

func (a *app) routes() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/", a.Home)
	mux.HandleFunc("/quests", a.Quests)
	mux.HandleFunc("/quests/add", a.AddQuest)

	return mux
}
