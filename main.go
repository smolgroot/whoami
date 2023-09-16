package main

import (
	"net/http"
	"os"
	"text/template"
)

type PageData struct {
	PageTitle string
	CurrentIp string
}

var tpl = template.Must(template.ParseFiles("index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		PageTitle: "NETBLIX",
		CurrentIp: "10.0.0.1",
	}
	tpl.Execute(w, data)
}

func main() {
	fs := http.FileServer(http.Dir("assets"))
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8088"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.ListenAndServe(":"+port, mux)
}
