package main

import (
	"net/http"
	"text/template"
)

type LoginPage struct {
	Labelone string
	Labeltwo string
	MsgError string
}

func LoginHtml(w http.ResponseWriter, req *http.Request) {
	tmpl := template.Must(template.ParseFiles("login.html"))
	data := LoginPage{
		Labelone: "Sign in",
		Labeltwo: "",
	}
	tmpl.Execute(w, data)
}

func Auth(w http.ResponseWriter, r *http.Request) {
	json := `{"Status":"ok", "Message":"tudo ocorreu bem na Auth..."}`
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(json))
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/login", LoginHtml)
	mux.HandleFunc("/v1/api/auth", Auth)

	// diretorio fisico
	fs := http.FileServer(http.Dir("./web/static"))

	// mostra no browser localhost:8080/static
	mux.Handle("/web/static/", http.StripPrefix("/web/static", fs))

	http.ListenAndServe(":8080", mux)
}
