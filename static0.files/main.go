package main

import "net/http"

func main() {

	// diretorio fisico
	fs := http.FileServer(http.Dir("assets/"))

	// mostra no browser localhost:8080/static
	//http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.Handle("/", http.StripPrefix("/", fs))

	// listener
	http.ListenAndServe(":8080", nil)
}
