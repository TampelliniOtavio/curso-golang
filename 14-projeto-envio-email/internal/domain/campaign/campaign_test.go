package campaign

import (
	"testing"
	"time"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

var (
    name = "Campaign X"
    createdBy = "teste@teste.com.br"
    content = "Content"
    contacts = []string{"email@gmail.com", "email2@gmail.com"}
    fake = faker.New()
)

func Test_NewCampaign_CreateNewCampaign(t *testing.T) {
    assert := assert.New(t)
    campaign, _ := NewCampaign(name, content, contacts, createdBy)

    assert.Equal(campaign.Name, name)
    assert.Equal(campaign.Content, content)
    assert.Equal(campaign.Status, Pending)
    assert.Equal(len(campaign.Contacts), len(contacts))
    assert.Equal(campaign.CreatedBy, createdBy)
}

func Test_NewCampaign_IdIsNotNil(t *testing.T) {
    assert := assert.New(t)

    campaign, _ := NewCampaign(name, content, contacts, createdBy)

    assert.NotNil(campaign.Id)
}

func Test_NewCampaign_CreatedOnMustBeNow(t *testing.T) {
    assert := assert.New(t)
    now := time.Now().Add(-time.Minute)

    campaign, _ := NewCampaign(name, content, contacts, createdBy)

    assert.Greater(campaign.CreatedOn, now)
}

func Test_NewCampaign_MustValidateNameMin(t *testing.T) {
    assert := assert.New(t)

    _, err := NewCampaign(fake.Lorem().Text(4), content, contacts, createdBy)

    assert.Equal("name is required with min 5", err.Error())
}

func Test_NewCampaign_MustValidateNameMax(t *testing.T) {
    assert := assert.New(t)

    _, err := NewCampaign(fake.Lorem().Text(25), content, contacts, createdBy)

    assert.Equal("name is required with max 24", err.Error())
}

func Test_NewCampaign_MustValidateContentMin(t *testing.T) {
    assert := assert.New(t)

    _, err := NewCampaign(name, fake.Lorem().Text(4), contacts, createdBy)

    assert.Equal("content is required with min 5", err.Error())
}

func Test_NewCampaign_MustValidateContentMax(t *testing.T) {
    assert := assert.New(t)

    _, err := NewCampaign(name, fake.Lorem().Text(1040), contacts, createdBy)

    assert.Equal("content is required with max 1024", err.Error())
}

func Test_NewCampaign_MustValidateContacts(t *testing.T) {
    assert := assert.New(t)

    _, err := NewCampaign(name, content, make([]string, 0), createdBy)

    assert.Equal("contacts is required with min 1", err.Error())
}

func Test_NewCampaign_MustValidateContactsEmailInvalid(t *testing.T) {
    assert := assert.New(t)

    _, err := NewCampaign(name, content, []string{""}, createdBy)

    assert.Equal("email is invalid", err.Error())
}

func Test_NewCampaign_MustValidateCreatedBy(t *testing.T) {
    assert := assert.New(t)

    _, err := NewCampaign(name, content, contacts, "")

    assert.Equal("createdby is required", err.Error())
}
