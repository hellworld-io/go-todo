package main

import (
	"net/http"
	"text/template"
)

func mainPage(w http.ResponseWriter, r *http.Request) {
	mainTemplate, _ := template.ParseFiles("/Users/yoonjeongbu/coma/dev/go-todo/www/html/todo_main.html")
	mainTemplate.Execute(w, nil)
}

func main() {
	//	fmt.Println("Hello")
	http.Handle("/www/", http.StripPrefix("/www/", http.FileServer(http.Dir("."))))
	http.HandleFunc("/", mainPage)
	http.ListenAndServe(":8080", nil)
}
