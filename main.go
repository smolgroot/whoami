package main

import (
	"fmt"
	"net"
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
	host, _, _ := net.SplitHostPort(r.RemoteAddr)
	fmt.Sprintln(host)
	data := PageData{
		PageTitle: "NETBLIX",
		CurrentIp: host,
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
