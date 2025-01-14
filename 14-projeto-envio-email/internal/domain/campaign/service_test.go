package campaign_test

import (
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
	newCampaign = campaign.NewCampaignRequest{
		Name:      "Super Name",
		Content:   "Content",
		Emails:    []string{"email1@email.com"},
		CreatedBy: "email@email.com",
	}
	campaignPending *campaign.Campaign
	campaignStarted *campaign.Campaign
	repository      = new(internalmock.CampaignRepositoryMock)
	service         = campaign.ServiceImp{
		Repository: repository,
	}
)

func setUp() {
	repository = new(internalmock.CampaignRepositoryMock)
	service = campaign.ServiceImp{
		Repository: repository,
	}
	campaignPending, _ = campaign.NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails, newCampaign.CreatedBy)
	campaignStarted, _ = campaign.NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails, newCampaign.CreatedBy)
	campaignStarted.Status = campaign.Started
}

func setUpGetById(campaign *campaign.Campaign, err error) {
	repository.On("GetById", mock.Anything).Return(campaign, err)
}

func setUpUpdate(err error) {
	repository.On("Update", mock.Anything).Return(err)
}

func setUpEmailSuccess() {
	service.SendMail = func(campaign *campaign.Campaign) error {
		return nil
	}
}

func setUpEmailError(err error) {
	service.SendMail = func(campaign *campaign.Campaign) error {
		return err
	}
}

func Test_Create_ValidateDomainError(t *testing.T) {
	setUp()
	errorCampaign := campaign.NewCampaignRequest{
		Name:    "",
		Content: "Body",
		Emails:  []string{"email1@email.com"},
	}
	assert := assert.New(t)

	_, err := service.Create(errorCampaign)

	assert.NotNil(err)
	assert.False(errors.Is(internalerrors.ErrInternal, err))
}

func Test_Create_Campaign(t *testing.T) {
	setUp()
	assert := assert.New(t)

	repository.On("Create", mock.Anything).Return(nil)

	id, err := service.Create(newCampaign)

	assert.NotNil(id)
	assert.Nil(err)
}

func Test_Create_ValidateRepositorySave(t *testing.T) {
	setUp()
	assert := assert.New(t)

	repository.On("Create", mock.Anything).Return(internalerrors.ErrInternal)

	_, err := service.Create(newCampaign)

	assert.True(errors.Is(internalerrors.ErrInternal, err))
}

func Test_GetById_ReturnCampaign(t *testing.T) {
	setUp()
	assert := assert.New(t)

	repository.On("GetById", mock.MatchedBy(func(id string) bool {
		return id == campaignPending.Id
	})).Return(campaignPending, nil)

	campaignPending, err := service.GetById(campaignPending.Id)

	assert.NotNil(campaignPending)
	assert.Nil(err)
	assert.Equal(campaignPending.Id, campaignPending.Id)
	assert.Equal(campaignPending.Status, campaignPending.Status)
	assert.Equal(campaignPending.Name, campaignPending.Name)
	assert.Equal(campaignPending.Content, campaignPending.Content)
	assert.Equal(campaignPending.CreatedBy, campaignPending.CreatedBy)
}

func Test_GetById_ReturnError(t *testing.T) {
	setUp()
	assert := assert.New(t)

	setUpGetById(nil, internalerrors.ErrInternal)

	_, err := service.GetById(campaignPending.Id)
	assert.Equal(internalerrors.ErrInternal.Error(), err.Error())
}

func Test_Delete_ReturnRecordNotFound(t *testing.T) {
	setUp()
	assert := assert.New(t)

	setUpGetById(nil, gorm.ErrRecordNotFound)

	err := service.Delete("invalid_id")

	assert.NotNil(err)
	assert.True(errors.Is(err, gorm.ErrRecordNotFound))
}

func Test_Delete_CampaignStatusInvalid(t *testing.T) {
	setUp()
	assert := assert.New(t)

	setUpGetById(campaignStarted, nil)

	err := service.Delete(campaignStarted.Id)

	assert.NotNil(err)
	assert.Equal(err.Error(), "Campaign status Invalid")
}

func Test_Delete_InternalError_when_something_goes_wrong(t *testing.T) {
	setUp()
	assert := assert.New(t)

	setUpGetById(campaignPending, nil)
	repository.On("Delete", mock.Anything).Return(errors.New("error to delete campaign"))

	err := service.Delete(campaignPending.Id)

	assert.NotNil(err)
	assert.Equal(err.Error(), internalerrors.ErrInternal.Error())
}

func Test_Delete_Success(t *testing.T) {
	setUp()
	assert := assert.New(t)

	setUpGetById(campaignPending, nil)
	repository.On("Delete", mock.MatchedBy(func(campaign *campaign.Campaign) bool {
		return campaignPending.Id == campaign.Id
	})).Return(nil)

	err := service.Delete(campaignPending.Id)

	assert.Nil(err)
}

func Test_Start_ReturnRecordNotFound(t *testing.T) {
	setUp()
	assert := assert.New(t)

	setUpGetById(nil, gorm.ErrRecordNotFound)

	err := service.Start("id_not_found")

	assert.NotNil(err)
	assert.True(errors.Is(err, gorm.ErrRecordNotFound))
}

func Test_Start_CampaignNotPending(t *testing.T) {
	setUp()
	assert := assert.New(t)

	setUpGetById(campaignStarted, nil)

	err := service.Start(campaignStarted.Id)

	assert.NotNil(err)
	assert.Equal("Campaign status Invalid", err.Error())
}

func Test_Start_ShouldSendMail_UpdateToStarted(t *testing.T) {
	setUp()
	assert := assert.New(t)

	setUpGetById(campaignPending, nil)

	repository.On("Update", mock.MatchedBy(func(campaignUpd *campaign.Campaign) bool {
		return campaignPending.Id == campaignUpd.Id && campaignUpd.Status == campaign.Started
	})).Return(nil)

	setUpEmailSuccess()

	service.Start(campaignPending.Id)

	assert.Equal(campaign.Started, campaignPending.Status)
}

func Test_SendEmailAndUpdateStatus_OnFailSetStatusFail(t *testing.T) {
	setUp()

	setUpEmailError(errors.New("An Error Ocourred"))
	repository.On("Update", mock.MatchedBy(func(campaignUpd *campaign.Campaign) bool {
		return campaignPending.Id == campaignUpd.Id && campaignUpd.Status == campaign.Failed
	})).Return(nil)

	service.SendEmailAndUpdateStatus(campaignPending)

	repository.AssertExpectations(t)
}

func Test_SendEmailAndUpdateStatus_OnSuccessSetStatusDone(t *testing.T) {
	setUp()

	setUpEmailSuccess()
	repository.On("Update", mock.MatchedBy(func(campaignUpd *campaign.Campaign) bool {
		return campaignPending.Id == campaignUpd.Id && campaignUpd.Status == campaign.Done
	})).Return(nil)

	service.SendEmailAndUpdateStatus(campaignPending)

	repository.AssertExpectations(t)
}
