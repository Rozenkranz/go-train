package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

var InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
var ErrorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

type app struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
}

func main() {

	addr := flag.String("addr", "localhost:8080", "Address to host HTTP server")

	flag.Parse()

	a := &app{
		ErrorLog: ErrorLog,
		InfoLog:  InfoLog,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: ErrorLog,
		Handler:  a.routes(),
	}

	InfoLog.Printf("Starting server on %v...\n", *addr)
	err := srv.ListenAndServe()
	ErrorLog.Fatal(err)

}
