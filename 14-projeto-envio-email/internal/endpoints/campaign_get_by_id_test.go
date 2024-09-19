package endpoints

import (
	"emailn/internal/contract"
	"emailn/internal/domain/campaign"
	internalMock "emailn/internal/test/mock"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CampaignGetById_should_return_campaign(t *testing.T) {
    assert := assert.New(t)

    body := contract.CampaignResponse{
        Id: "ID",
        Name: "Name",
        Status: campaign.Pending,
        Content: "Content",
    }

    service := &internalMock.CampaignServiceMock{}
    service.On("GetById", mock.Anything).Return(&body, nil)
    handler := Handler{CampaignService: service}

    req, _ := http.NewRequest("GET", "/campaigns/" + body.Id, nil)
    rr := httptest.NewRecorder()

    response, status, err := handler.CampaignGetById(rr, req)

    campaign := response.(*contract.CampaignResponse)

    assert.Nil(err)
    assert.Equal(status, 200)
    assert.Equal(campaign.Id, body.Id)
    assert.Equal(campaign.Name, body.Name)
    assert.Equal(campaign.Content, body.Content)
    assert.Equal(campaign.Status, body.Status)
}

func Test_CampaignGetById_should_return_error(t *testing.T) {
    assert := assert.New(t)

    errExpected := errors.New("Something went Wrong")

    service := &internalMock.CampaignServiceMock{}
    service.On("GetById", mock.Anything).Return(nil, errExpected)
    handler := Handler{CampaignService: service}

    req, _ := http.NewRequest("GET", "/", nil)
    rr := httptest.NewRecorder()

    _, status, err := handler.CampaignGetById(rr, req)

    assert.NotNil(err)
    assert.Equal(err.Error(), errExpected.Error())
    assert.Equal(status, 200)
}
