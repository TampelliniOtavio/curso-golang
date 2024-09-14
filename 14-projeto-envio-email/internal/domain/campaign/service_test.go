package campaign

import (
	"emailn/internal/contract"
	internalerrors "emailn/internal/internal-errors"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct{
    mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error {
    args := r.Called(campaign)
    return args.Error(0)
}

var (
    service = Service{}
)

func Test_Create_ValidateDomainError(t *testing.T) {
    newCampaign := contract.NewCampaign{
        Name: "",
        Content: "Body",
        Emails: []string{"email1@email.com"},
    }
    assert := assert.New(t)

    _, err := service.Create(newCampaign)

    assert.NotNil(err)
    assert.Equal("name is required", err.Error())
}

func Test_Create_Campaign(t *testing.T) {
    newCampaign := contract.NewCampaign{
        Name: "Test",
        Content: "Body",
        Emails: []string{"email1@email.com"},
    }

    repository := new(repositoryMock)
    repository.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
        if newCampaign.Name != campaign.Name {
            return false
        }

        if newCampaign.Content != campaign.Content {
            return false
        }

        if len(newCampaign.Emails) != len(campaign.Contacts) {
            return false
        }

        return true
    })).Return(nil)

    service.Repository = repository

    service.Create(newCampaign)

    repository.AssertExpectations(t)
}

func Test_Create_ValidateRepositorySave(t *testing.T) {
    assert := assert.New(t)
    newCampaign := contract.NewCampaign{
        Name: "Test",
        Content: "Body",
        Emails: []string{"email1@email.com"},
    }

    repository := new(repositoryMock)
    repository.On("Save", mock.Anything).Return(errors.New("error to save on database"))
    service.Repository = repository

    _, err := service.Create(newCampaign)

    assert.True(errors.Is(internalerrors.ErrInternal, err))
}
