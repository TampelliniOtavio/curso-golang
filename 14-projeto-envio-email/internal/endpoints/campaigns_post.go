package endpoints

import (
	"emailn/internal/contract"
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) CampaignsPost (w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
    var request contract.NewCampaign
    render.DecodeJSON(r.Body, &request)
    request.CreatedBy = r.Context().Value("email").(string)
    id, err := h.CampaignService.Create(request)
    return map[string]string{"id": id}, 201, err
}
