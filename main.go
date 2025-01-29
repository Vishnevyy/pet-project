package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var task string

func GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if task == "" {
			fmt.Fprintln(w, "hello, no task set")
		} else {
			fmt.Fprintf(w, "hello, %s\n", task)
		}
	} else {
		fmt.Fprintln(w, "Поддерживается только метод Get")
	}
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var data map[string]string

		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if t, ok := data["task"]; ok {
			task = t
			fmt.Fprintf(w, "Task сохранён: %s\n", task)
		} else {
			http.Error(w, "Поле 'task' обязательно.", http.StatusBadRequest)
		}
	} else {
		fmt.Fprintln(w, "Поддерживается только метод Post")
	}
}

func main() {
	http.HandleFunc("/get", GetHandler)
	http.HandleFunc("/post", PostHandler)

	http.ListenAndServe("localhost:8080", nil)
}
