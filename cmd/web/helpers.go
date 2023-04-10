package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// Выводит ошибку в лог
func (a *app) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	a.ErrorLog.Println(trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// Выводит ошибку пользователю
func (a *app) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// Если страница не найдена
func (a *app) notFound(w http.ResponseWriter) {
	a.clientError(w, http.StatusNotFound)
}
