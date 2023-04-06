package main

import (
	"flag"
	"net/http"

	"github.com/Rozenkranz/go-train/cmd/web/handlers"
	"github.com/Rozenkranz/go-train/logging"
)

func main() {

	addr := flag.String("addr", "localhost:8080", "Address to host HTTP server")

	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/quests", handlers.Quests)
	mux.HandleFunc("/quests/add", handlers.AddQuest)

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: logging.ErrorLog,
		Handler:  mux,
	}

	logging.InfoLog.Printf("Starting server on %v...\n", *addr)
	err := srv.ListenAndServe()
	logging.ErrorLog.Fatal(err)

}
