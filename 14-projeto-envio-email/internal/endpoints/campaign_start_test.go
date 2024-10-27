package endpoints

import (
	"context"
	"emailn/internal/test/internalmock"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CampaignStart_204(t *testing.T) {
    assert := assert.New(t)

    service := &internalmock.CampaignServiceMock{}
    service.On("Start", mock.Anything).Return(nil)
    handler := Handler{CampaignService: service}

    campaignId := "id_valid"

    service.On("Start", mock.MatchedBy(func (id string) bool {
        return campaignId == id
    })).Return(nil)

    req, _ := http.NewRequest("PATCH", "/campaigns/start/" + campaignId, nil)

    chiContext := chi.NewRouteContext()
    chiContext.URLParams.Add("id", campaignId) 

    req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiContext))

    rr := httptest.NewRecorder()

    _, status, err := handler.CampaignStart(rr, req)

    assert.Nil(err)
    assert.Equal(status, 204)
}

func Test_CampaignStart_500(t *testing.T) {
    assert := assert.New(t)

    errExpected := errors.New("Something went Wrong")

    service := &internalmock.CampaignServiceMock{}
    service.On("Start", mock.Anything).Return(errExpected)
    handler := Handler{CampaignService: service}

    req, _ := http.NewRequest("PATCH", "/campaigns/start/id_invalid", nil)
    rr := httptest.NewRecorder()

    _, _, err := handler.CampaignStart(rr, req)

    assert.NotNil(err)
    assert.Equal(err.Error(), errExpected.Error())
}
