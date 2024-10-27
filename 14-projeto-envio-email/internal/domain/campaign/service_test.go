package campaign_test

import (
	"emailn/internal/contract"
	"emailn/internal/domain/campaign"
	internalerrors "emailn/internal/internal-errors"
	"emailn/internal/test/internalmock"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var (
    service = campaign.ServiceImp{}
    newCampaign = contract.NewCampaign{
        Name: "Super Name",
        Content: "Content",
        Emails: []string{"email1@email.com"},
        CreatedBy: "email@email.com",
    }
)

func Test_Create_ValidateDomainError(t *testing.T) {
    errorCampaign := contract.NewCampaign{
        Name: "",
        Content: "Body",
        Emails: []string{"email1@email.com"},
    }
    assert := assert.New(t)

    _, err := service.Create(errorCampaign)

    assert.NotNil(err)
    assert.False(errors.Is(internalerrors.ErrInternal, err))
}

func Test_Create_Campaign(t *testing.T) {
    assert := assert.New(t)

    repository := new(internalmock.CampaignRepositoryMock)
    repository.On("Create", mock.Anything).Return(nil)
    service.Repository = repository

    id, err := service.Create(newCampaign)

    assert.NotNil(id)
    assert.Nil(err)
}

func Test_Create_ValidateRepositorySave(t *testing.T) {
    assert := assert.New(t)

    repository := new(internalmock.CampaignRepositoryMock)
    repository.On("Create", mock.Anything).Return(internalerrors.ErrInternal)
    service.Repository = repository

    _, err := service.Create(newCampaign)

    assert.True(errors.Is(internalerrors.ErrInternal, err))
}

func Test_GetById_ReturnCampaign(t *testing.T) {
    assert := assert.New(t)

    campaign, _ := campaign.NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails, newCampaign.CreatedBy)

    repository := new(internalmock.CampaignRepositoryMock)
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
    assert.Equal(campaign.CreatedBy, campaignReturned.CreatedBy)
}

func Test_GetById_ReturnError(t *testing.T) {
    assert := assert.New(t)

    campaign, _ := campaign.NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails, newCampaign.CreatedBy)

    repository := new(internalmock.CampaignRepositoryMock)
    repository.On("GetById", mock.Anything).Return(nil, internalerrors.ErrInternal)
    service.Repository = repository

    _, err := service.GetById(campaign.Id)
    assert.Equal(internalerrors.ErrInternal.Error(), err.Error())
}

func Test_Delete_ReturnRecordNotFound(t *testing.T) {
    assert := assert.New(t)

    campaign, _ := campaign.NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails, newCampaign.CreatedBy)

    repository := new(internalmock.CampaignRepositoryMock)
    repository.On("GetById", mock.Anything).Return(nil, gorm.ErrRecordNotFound)
    service.Repository = repository

    err := service.Delete(campaign.Id)

    assert.NotNil(err)
    assert.True(errors.Is(err, gorm.ErrRecordNotFound))
}

func Test_Delete_CampaignStatusInvalid(t *testing.T) {
    assert := assert.New(t)
    campaignReturned, _ := campaign.NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails, newCampaign.CreatedBy)
    campaignReturned.Status = campaign.Started

    repository := new(internalmock.CampaignRepositoryMock)
    repository.On("GetById", mock.Anything).Return(campaignReturned, nil)
    service.Repository = repository

    err := service.Delete(campaignReturned.Id)

    assert.NotNil(err)
    assert.Equal(err.Error(), "Campaign status Invalid")
}

func Test_Delete_InternalError_when_something_goes_wrong(t *testing.T) {
    assert := assert.New(t)
    campaignReturned, _ := campaign.NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails, newCampaign.CreatedBy)

    repository := new(internalmock.CampaignRepositoryMock)
    repository.On("GetById", mock.Anything).Return(campaignReturned, nil)
    repository.On("Delete", mock.MatchedBy(func (campaign *campaign.Campaign) bool {
        return campaignReturned.Id == campaign.Id
    })).Return(errors.New("error to delete campaign"))
    service.Repository = repository

    err := service.Delete(campaignReturned.Id)

    assert.NotNil(err)
    assert.Equal(err.Error(), internalerrors.ErrInternal.Error())
}

func Test_Delete_Success(t *testing.T) {
    assert := assert.New(t)
    campaignReturned, _ := campaign.NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails, newCampaign.CreatedBy)

    repository := new(internalmock.CampaignRepositoryMock)
    repository.On("GetById", mock.Anything).Return(campaignReturned, nil)
    repository.On("Delete", mock.MatchedBy(func (campaign *campaign.Campaign) bool {
        return campaignReturned.Id == campaign.Id
    })).Return(nil)
    service.Repository = repository

    err := service.Delete(campaignReturned.Id)

    assert.Nil(err)
}

func Test_Start_ReturnRecordNotFound(t *testing.T) {
    assert := assert.New(t)

    campaign, _ := campaign.NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails, newCampaign.CreatedBy)

    repository := new(internalmock.CampaignRepositoryMock)
    repository.On("GetById", mock.Anything).Return(nil, gorm.ErrRecordNotFound)
    service.Repository = repository

    err := service.Start(campaign.Id)

    assert.NotNil(err)
    assert.True(errors.Is(err, gorm.ErrRecordNotFound))
}

func Test_Start_CampaignNotPending(t *testing.T) {
    assert := assert.New(t)

    campaignCreated, _ := campaign.NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails, newCampaign.CreatedBy)
    campaignCreated.Status = campaign.Done

    repository := new(internalmock.CampaignRepositoryMock)
    repository.On("GetById", mock.Anything).Return(campaignCreated, nil)
    service.Repository = repository

    err := service.Start(campaignCreated.Id)

    assert.NotNil(err)
    assert.Equal("Campaign status Invalid", err.Error())
}

func Test_Start_ShouldSendMail(t *testing.T) {
    assert := assert.New(t)

    campaignCreated, _ := campaign.NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails, newCampaign.CreatedBy)

    repository := new(internalmock.CampaignRepositoryMock)
    repository.On("GetById", mock.Anything).Return(campaignCreated, nil)
    repository.On("Update", mock.Anything).Return(nil)
    service.Repository = repository
    mailWasSent := false
    sendMail := func(campaign *campaign.Campaign) error {
        if campaign.Id == campaignCreated.Id {
            mailWasSent = true
        }
        return nil
    }
    service.SendMail = sendMail

    err := service.Start(campaignCreated.Id)

    assert.Nil(err)
    assert.True(mailWasSent)
}

func Test_Start_ReturnError_fail(t *testing.T) {
    assert := assert.New(t)

    campaignCreated, _ := campaign.NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails, newCampaign.CreatedBy)

    repository := new(internalmock.CampaignRepositoryMock)
    repository.On("GetById", mock.Anything).Return(campaignCreated, nil)
    service.Repository = repository
    sendMail := func(campaign *campaign.Campaign) error {
        return errors.New("error to send mail")
    }
    service.SendMail = sendMail

    err := service.Start(campaignCreated.Id)

    assert.NotNil(err)
    assert.True(errors.Is(internalerrors.ErrInternal, err))
}

func Test_Start_ShouldSendMail_UpdateToDone(t *testing.T) {
    assert := assert.New(t)

    campaignCreated := &campaign.Campaign{ Id: "1", Status: campaign.Pending}

    repository := new(internalmock.CampaignRepositoryMock)
    repository.On("GetById", mock.Anything).Return(campaignCreated, nil)
    repository.On("Update", mock.MatchedBy(func (campaignUpd *campaign.Campaign) bool {
        return campaignCreated.Id == campaignUpd.Id && campaignUpd.Status == campaign.Done
    })).Return(nil)

    service.Repository = repository
    sendMail := func(campaign *campaign.Campaign) error {
        return nil
    }
    service.SendMail = sendMail

    service.Start(campaignCreated.Id)

    assert.Equal(campaign.Done, campaignCreated.Status)
}
