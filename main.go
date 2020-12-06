package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	destination := vars["destination"]

	layoutTempl(destination).ExecuteTemplate(w, "layout_project", &Page{Title: strings.ToUpper(destination)})
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	templ("main").ExecuteTemplate(w, "layout", &Page{Title: "Welcome to Tipesss"})
}

func projectsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	project := vars["project"]
	projectTempl(project).ExecuteTemplate(w, "layout_project", &Page{Title: project})
}

func projectTempl(name string) *template.Template {
	t := template.New("")
	fmt.Println("templates/projects/" + name + ".html")
	temp, err := t.ParseFiles("templates/layout_project.html", "templates/header.html", "templates/projects/"+name+".html")

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return temp
}

func layoutTempl(name string) *template.Template {
	t := template.New("")
	fmt.Println("templates/projects/" + name + ".html")
	temp, err := t.ParseFiles("templates/layout_project.html", "templates/header.html", "templates/"+name+".html")

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return temp
}

func templ(name string) *template.Template {
	t := template.New("")
	fmt.Println("templates/projects/" + name + ".html")
	temp, err := t.ParseFiles("templates/layout.html", "templates/header.html", "templates/"+name+".html")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return temp
}

func main() {
	sfs := http.FileServer(http.Dir("./static"))
	m := mux.NewRouter()
	s := &http.Server{
		Addr:         ":8080",
		Handler:      m,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	m.PathPrefix("/static/").Handler(http.StripPrefix("/static/", sfs))

	m.HandleFunc("/", homeHandler)
	m.HandleFunc("/{destination}/", mainHandler)
	m.HandleFunc("/project/{project}/", projectsHandler)
	m.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) { s.Shutdown(context.Background()) })

	fmt.Println("Server Started on port 8080")
	log.Fatal(s.ListenAndServe())
}
