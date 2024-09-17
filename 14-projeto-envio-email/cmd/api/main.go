package main

import (
	"emailn/internal/domain/campaign"
	"emailn/internal/domain/infrastructure/database"
	"emailn/internal/endpoints"
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
        Repository: &database.CampaignRepository{},
    }

    handler := endpoints.Handler{
        CampaignService: &campaignSservice,
    }

    r.Post("/campaigns", endpoints.HandlerError(handler.CampaignsPost))
    r.Get("/campaigns", endpoints.HandlerError(handler.CampaignsGet))

    http.ListenAndServe(":3000", r)
}
