package endpoints

import (
	"emailn/internal/domain/campaign"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CampaignGetById_should_return_campaign(t *testing.T) {
    setup()
    assert := assert.New(t)

    body := campaign.CampaignResponse{
        Id: "ID",
        Name: "Name",
        Status: campaign.Pending,
        Content: "Content",
    }

    service.On("GetById", mock.Anything).Return(&body, nil)

    req, rr := newReqAndRecord("GET", "/", nil)

    response, status, err := handler.CampaignGetById(rr, req)

    campaign := response.(*campaign.CampaignResponse)

    assert.Nil(err)
    assert.Equal(status, 200)
    assert.Equal(campaign.Id, body.Id)
    assert.Equal(campaign.Name, body.Name)
    assert.Equal(campaign.Content, body.Content)
    assert.Equal(campaign.Status, body.Status)
}

func Test_CampaignGetById_should_return_error(t *testing.T) {
    setup()
    assert := assert.New(t)

    errExpected := errors.New("Something went Wrong")

    service.On("GetById", mock.Anything).Return(nil, errExpected)

    req, rr := newReqAndRecord("GET", "/", nil)

    _, status, err := handler.CampaignGetById(rr, req)

    assert.NotNil(err)
    assert.Equal(err.Error(), errExpected.Error())
    assert.Equal(status, 200)
}
