package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Todo struct {
	Id   string
	Name string
}

var todos []Todo

func main() {
	t, e := template.ParseGlob("*.html")
	if e != nil {
		log.Panic(e)
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			t.ExecuteTemplate(w, "index.html", todos)
		case http.MethodPost:
			todos = append(todos, Todo{
				Id:   strconv.Itoa(len(todos)),
				Name: r.FormValue("name"),
			})
		case http.MethodPatch:
			for i, v := range todos {
				if v.Id == r.FormValue("id") {
					todos[i].Name = r.FormValue("name")
					return
				}
			}
		case http.MethodDelete:
			var out []Todo
			for _, v := range todos {
				if v.Id != r.FormValue("id") {
					out = append(out, v)
				}
			}
		}
	})
	http.ListenAndServe(":8080", nil)
}
