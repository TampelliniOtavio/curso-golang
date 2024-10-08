package endpoints

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) CampaignGetById(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
    id := chi.URLParam(r, "id")

    campaign, err := h.CampaignService.GetById(id)

    return campaign, 200, err
}
