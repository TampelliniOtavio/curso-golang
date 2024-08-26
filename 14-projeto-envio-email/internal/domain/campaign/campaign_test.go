package campaign

import "testing"

func TestNewCampaign(t *testing.T) {
    name := "Campaign X"
    content := "Body"
    contacts := []string{"email@gmail.com", "email2@gmail.com"}

    campaign := NewCampaign(name, content, contacts)

    if campaign.Id != "1" {
        t.Errorf("expected 1")
    } else if campaign.Name != name {
        t.Errorf("expected correct name")
    } else if campaign.Content != content {
        t.Errorf("expected correct content")
    }

}
