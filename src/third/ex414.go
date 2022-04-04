package main

import (
	"html/template"
	"learn/src/third/git"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe("localhost:8888", nil))
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println(err)
	}
	q := []string{r.FormValue("key")}
	result, err := git.SearchIssues(q)
	if err != nil {
		log.Println(err)
	}
	tmpl := template.Must(template.ParseFiles("index.html"))
	if err := tmpl.Execute(w, result); err != nil {
		log.Println(err)
	}
}
