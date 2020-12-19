package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/ping", Ping)
	log.Println("Run Server 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func Ping(w http.ResponseWriter, r *http.Request) {
	log.Println("server ok")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"msg":"pong"}`))
}
