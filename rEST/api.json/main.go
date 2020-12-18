package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type StUser struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/user", User)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func User(w http.ResponseWriter, r *http.Request) {

	var u StUser
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error:", err.Error())
		w.Write([]byte(`{"msg":"Error newDecoder"}`))
		return
	}

	name := u.Name

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"server.name:"` + name + `"}`))
}
