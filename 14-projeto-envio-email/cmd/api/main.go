package main

import (
	"emailn/internal/domain/campaign"
	"emailn/internal/endpoints"
	"emailn/internal/infrastructure/database"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
    r := chi.NewRouter()

    r.Use(middleware.RequestID)
    r.Use(middleware.RealIP)
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)

    campaignSservice := campaign.ServiceImp{
        Repository: &database.CampaignRepository{
            Db: database.NewDB(),
        },
    }

    handler := endpoints.Handler{
        CampaignService: &campaignSservice,
    }

    r.Route("/campaigns", func (r chi.Router) {
        r.Use(endpoints.Auth)
        r.Post("/", endpoints.HandlerError(handler.CampaignsPost))
        r.Get("/", endpoints.HandlerError(handler.CampaignsGet))
        r.Get("/{id}", endpoints.HandlerError(handler.CampaignGetById))
        r.Delete("/{id}", endpoints.HandlerError(handler.CampaignDelete))
    })


    print("Server Starting...\n")
    http.ListenAndServe(":3000", r)
}
