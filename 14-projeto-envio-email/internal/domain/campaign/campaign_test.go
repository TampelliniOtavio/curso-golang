package campaign

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

var (
    name = "Campaign X"
    content = "Body"
    contacts = []string{"email@gmail.com", "email2@gmail.com"}
)

func Test_NewCampaign_CreateNewCampaign(t *testing.T) {
    assert := assert.New(t)
    campaign, _ := NewCampaign(name, content, contacts)

    assert.Equal(campaign.Name, name)
    assert.Equal(campaign.Content, content)
    assert.Equal(len(campaign.Contacts), len(contacts))
}

func Test_NewCampaign_IdIsNotNil(t *testing.T) {
    assert := assert.New(t)

    campaign, _ := NewCampaign(name, content, contacts)

    assert.NotNil(campaign.Id)
}

func Test_NewCampaign_CreatedOnMustBeNow(t *testing.T) {
    assert := assert.New(t)
    now := time.Now().Add(-time.Minute)

    campaign, _ := NewCampaign(name, content, contacts)

    assert.Greater(campaign.CreatedOn, now)
}

func Test_NewCampaign_MustValidateName(t *testing.T) {
    assert := assert.New(t)

    _, err := NewCampaign("", content, contacts)

    assert.Equal("name is required", err.Error())
}

func Test_NewCampaign_MustValidateContent(t *testing.T) {
    assert := assert.New(t)

    _, err := NewCampaign(name, "", contacts)

    assert.Equal("content is required", err.Error())
}

func Test_NewCampaign_MustValidateContacts(t *testing.T) {
    assert := assert.New(t)

    _, err := NewCampaign(name, content, make([]string, 0))

    assert.Equal("emails is required", err.Error())
}
