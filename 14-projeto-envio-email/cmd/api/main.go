package main

import (
	"emailn/internal/contract"
	"emailn/internal/domain/campaign"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func main() {
    r := chi.NewRouter()

    r.Use(middleware.RequestID)
    r.Use(middleware.RealIP)
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)

    service := campaign.Service{}

    r.Post("/campaigns", func(w http.ResponseWriter, r *http.Request) {
        var request contract.NewCampaign
        err := render.DecodeJSON(r.Body, &request)

        if err != nil {
            render.Status(r, 400)
            render.JSON(w, r, map[string]string{"error": err.Error()})
            return
        }

        id, err := service.Create(request)

        if err != nil {
            render.Status(r, 400)
            render.JSON(w, r, map[string]string{"error": err.Error()})
            return
        }

        render.Status(r, 201)
        render.JSON(w, r, map[string]string{"id": id})
    })

    http.ListenAndServe(":3000", r)
}
