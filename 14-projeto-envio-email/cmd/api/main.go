package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type product struct{
    ID int
    Name string
}

type myHandler struct{}
func (m myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("MyHandler"))
}

func main() {
    r := chi.NewRouter()

    m := myHandler{}
    r.Handle("/handler", m)

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

    r.Get("/json", func(w http.ResponseWriter, r *http.Request) {
        obj := map[string]string{"message": "success"}

        render.JSON(w, r, obj)
    })

    r.Post("/product", func(w http.ResponseWriter, r *http.Request) {
        var product product
        render.DecodeJSON(r.Body, &product)
        product.ID = 5
        render.JSON(w, r, product)
    })

    r.Put("/product", func(w http.ResponseWriter, r *http.Request) {
        var product product
        render.DecodeJSON(r.Body, &product)
        product.ID = 5
        render.JSON(w, r, product)
    })

    http.ListenAndServe(":3000", r)
}
