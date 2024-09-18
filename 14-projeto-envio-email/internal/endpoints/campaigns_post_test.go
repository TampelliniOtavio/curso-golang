package endpoints

import (
	"bytes"
	"emailn/internal/contract"
	internalMock "emailn/internal/test/mock"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CampaignsPost_should_save_new_campaign(t *testing.T) {
    assert := assert.New(t)

    body := contract.NewCampaign{
        Name: "mouse",
        Content: "Content",
        Emails: []string{"email@email.com"},
    }
    service := &internalMock.CampaignServiceMock{}
    service.On("Create", mock.MatchedBy(func(request contract.NewCampaign) bool {
        if request.Name == body.Name && request.Content == body.Content {
            return true
        }

        return false
    })).Return("34x", nil)
    handler := Handler{
        CampaignService: service,
    }

    var buf bytes.Buffer
    json.NewEncoder(&buf).Encode(body)

    req, _ := http.NewRequest("POST", "/", &buf)

    rr := httptest.NewRecorder()

    _, status, err := handler.CampaignsPost(rr, req)

    assert.Equal(201, status)
    assert.Nil(err)
}

func Test_CampaignsPost_should_inform_error_when_exist(t *testing.T) {
    assert := assert.New(t)

    body := contract.NewCampaign{
        Name: "mouse",
        Content: "Content",
        Emails: []string{"email@email.com"},
    }
    service := &internalMock.CampaignServiceMock{}
    service.On("Create", mock.Anything).Return("", errors.New("error"))
    handler := Handler{
        CampaignService: service,
    }

    var buf bytes.Buffer
    json.NewEncoder(&buf).Encode(body)

    req, _ := http.NewRequest("POST", "/campaigns", &buf)

    rr := httptest.NewRecorder()

    _, _, err := handler.CampaignsPost(rr, req)

    assert.NotNil(err)
}
