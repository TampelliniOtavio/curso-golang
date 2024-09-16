package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
    r := chi.NewRouter()
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        name := r.URL.Query().Get("name")
        id := r.URL.Query().Get("id")
        if name != "" || id != "" {
            w.Write([]byte(name + id))
            return
        }

        w.Write([]byte("Teste"))
    })

    r.Get("/{productName}", func(w http.ResponseWriter, r *http.Request) {
        param := chi.URLParam(r, "productName")
        w.Write([]byte(param))
    })

    http.ListenAndServe(":3000", r)
}
