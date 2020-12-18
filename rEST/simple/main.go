package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/ping",
		func(w http.ResponseWriter, r *http.Request) {
			log.Println("ok")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("pong"))
		})

	log.Fatal(http.ListenAndServe(":8080", mux))

}
