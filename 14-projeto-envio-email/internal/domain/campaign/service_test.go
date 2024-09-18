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

func (r *repositoryMock) GetById(id string) (*Campaign, error) {
    args := r.Called(id)
    err := args.Error(1)

    if err != nil {
        return nil, err
    }

    return args.Get(0).(*Campaign), nil
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

func Test_GetById_ReturnCampaign(t *testing.T) {
    assert := assert.New(t)
    newCampaign := contract.NewCampaign{
        Name: "Campaign",
        Content: "Content",
        Emails: []string{"email1@email.com"},
    }

    campaign, _ := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)

    repository := new(repositoryMock)
    repository.On("GetById", mock.MatchedBy(func (id string) bool {
        return id == campaign.Id
    })).Return(campaign, nil)
    service.Repository = repository

    campaignReturned, err := service.GetById(campaign.Id)

    assert.NotNil(campaignReturned)
    assert.Nil(err)
    assert.Equal(campaign.Id, campaignReturned.Id)
    assert.Equal(campaign.Status, campaignReturned.Status)
    assert.Equal(campaign.Name, campaignReturned.Name)
    assert.Equal(campaign.Content, campaignReturned.Content)
}

func Test_GetById_ReturnError(t *testing.T) {
    assert := assert.New(t)
    newCampaign := contract.NewCampaign{
        Name: "Campaign",
        Content: "Content",
        Emails: []string{"email1@email.com"},
    }

    campaign, _ := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)

    repository := new(repositoryMock)
    repository.On("GetById", mock.Anything).Return(nil, errors.New("Something Wrong!"))
    service.Repository = repository

    _, err := service.GetById(campaign.Id)

    assert.NotNil(err)
    assert.Equal(internalerrors.ErrInternal.Error(), err.Error())
}
