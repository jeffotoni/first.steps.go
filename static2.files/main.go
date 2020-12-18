//go get github.com/rakyll/statik
package main

import (
	"log"
	"net/http"

	_ "github.com/jeffotoni/first.steps.go/static2.files/statik" // TODO: Replace with the absolute import path
	"github.com/rakyll/statik/fs"
)

func main() {

	mux := http.NewServeMux()

	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	// diretorio fisico
	fsnew := http.FileServer(statikFS)

	// mostra no browser localhost:8080/static
	mux.Handle("/", http.StripPrefix("/", fsnew))

	log.Println("Run Server")
	http.ListenAndServe(":8080", mux)
}

// func DisabledFs(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		if strings.HasSuffix(r.URL.Path, "/") {
// 			http.NotFound(w, r)
// 			return
// 		}
// 		next.ServeHTTP(w, r)
// 	})
// }
