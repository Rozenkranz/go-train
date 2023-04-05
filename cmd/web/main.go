package main

import (
	"log"
	"net/http"

	"github.com/Rozenkranz/go-train/cmd/web/handlers"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/quests", handlers.Quests)
	mux.HandleFunc("/quests/add", handlers.AddQuest)

	log.Println("Starting server...")
	err := http.ListenAndServe("localhost:8181", mux)
	log.Fatal(err)

}
