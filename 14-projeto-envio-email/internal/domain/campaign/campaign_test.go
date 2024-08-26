package campaign

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewCampaign(t *testing.T) {
    assert := assert.New(t)
    name := "Campaign X"
    content := "Body"
    contacts := []string{"email@gmail.com", "email2@gmail.com"}

    campaign := NewCampaign(name, content, contacts)

    assert.Equal(campaign.Id, "1")
    assert.Equal(campaign.Name, name)
    assert.Equal(campaign.Content, content)
    assert.Equal(len(campaign.Contacts), len(contacts))
}
