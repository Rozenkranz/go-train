package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/Rozenkranz/go-train/logging"
)

func Home(w http.ResponseWriter, r *http.Request) {
	logging.InfoLog.Printf("Home was visited by %s\n with %s method\n", r.RemoteAddr, r.Method)

	var files = []string{
		"./ui/html/home_page.tmpl",
		"./ui/html/base_layout.tmpl",
		"./ui/html/footer_partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		logging.ErrorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		logging.ErrorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func Quests(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "You wanna see quests with %d ID", id)

	logging.InfoLog.Printf("Quests was visited by %s\n with %s method\n", r.RemoteAddr, r.Method)
}

func AddQuest(w http.ResponseWriter, r *http.Request) {
	//Проверка на тип запроса
	if r.Method != http.MethodPost {
		//Запись кода ошибки
		w.WriteHeader(405)
		//Добавление заголовка разрешённого метода
		w.Header().Set("Allow", http.MethodPost)
		w.Write([]byte("Your request is not allowed\n"))
		logging.InfoLog.Printf("Add was visited by %s\n with %s method\n", r.RemoteAddr, r.Method)
		return
	}
	w.Write([]byte("Creating a new quest"))
	logging.InfoLog.Printf("AddQuest was visited by %s\n with %s method\n", r.RemoteAddr, r.Method)
}
