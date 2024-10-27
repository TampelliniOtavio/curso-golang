package endpoints

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CampaignStart_204(t *testing.T) {
    setup()
    assert := assert.New(t)

    service.On("Start", mock.Anything).Return(nil)

    campaignId := "id_valid"

    service.On("Start", mock.MatchedBy(func (id string) bool {
        return campaignId == id
    })).Return(nil)

    req, rr := newReqAndRecord("PATCH", "/", nil)

    addParameter(req, "id", campaignId)

    _, status, err := handler.CampaignStart(rr, req)

    assert.Nil(err)
    assert.Equal(status, 204)
}

func Test_CampaignStart_500(t *testing.T) {
    setup()
    assert := assert.New(t)

    errExpected := errors.New("Something went Wrong")

    service.On("Start", mock.Anything).Return(errExpected)

    req, rr := newReqAndRecord("PATCH", "/", nil)

    _, _, err := handler.CampaignStart(rr, req)

    assert.NotNil(err)
    assert.Equal(err.Error(), errExpected.Error())
}
