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

func (r *repositoryMock) Get() ([]Campaign, error) {
    // args := r.Called(campaign)
    return nil, nil
}

var (
    service = ServiceImp{}
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
    assert.False(errors.Is(internalerrors.ErrInternal, err))
}

func Test_Create_Campaign(t *testing.T) {
    assert := assert.New(t)
    newCampaign := contract.NewCampaign{
        Name: "Campaign",
        Content: "Content",
        Emails: []string{"email1@email.com"},
    }

    repository := new(repositoryMock)
    repository.On("Save", mock.Anything).Return(nil)

    service.Repository = repository

    id, err := service.Create(newCampaign)

    assert.NotNil(id)
    assert.Nil(err)
}

func Test_Create_ValidateRepositorySave(t *testing.T) {
    assert := assert.New(t)
    newCampaign := contract.NewCampaign{
        Name: "Campaign",
        Content: "Content",
        Emails: []string{"email1@email.com"},
    }

    repository := new(repositoryMock)
    repository.On("Save", mock.Anything).Return(internalerrors.ErrInternal)
    service.Repository = repository

    _, err := service.Create(newCampaign)

    assert.True(errors.Is(internalerrors.ErrInternal, err))
}
