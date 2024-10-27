package endpoints

import (
	"bytes"
	"context"
	"emailn/internal/test/internalmock"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/go-chi/chi/v5"
)

var (
    service = new(internalmock.CampaignServiceMock)
    handler = Handler{
        CampaignService: service,
    }
)

func setup() {
    service = new(internalmock.CampaignServiceMock)
    handler = Handler{
        CampaignService: service,
    }
}

func newReqAndRecord(method string, url string, body interface{}) (*http.Request, *httptest.ResponseRecorder) {
    var buf bytes.Buffer
    json.NewEncoder(&buf).Encode(body)
    req, _ := http.NewRequest(method, url, &buf)
    rr := httptest.NewRecorder()

    return req, rr
}

func addParameter(req *http.Request, keyParameter string, value string) {
    chiContext := chi.NewRouteContext()
    chiContext.URLParams.Add(keyParameter, value)
    req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiContext))
}

func addContext(req *http.Request, keyParameter string, value string) *http.Request {
    ctx := context.WithValue(req.Context(), keyParameter, value)
    return req.WithContext(ctx)
}
