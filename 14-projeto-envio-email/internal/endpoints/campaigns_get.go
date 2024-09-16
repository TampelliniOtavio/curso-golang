package endpoints

import (
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) CampaignsGet (w http.ResponseWriter, r *http.Request) {
    render.Status(r, 200)
    render.JSON(w, r, h.CampaignService.Repository.Get())
}
