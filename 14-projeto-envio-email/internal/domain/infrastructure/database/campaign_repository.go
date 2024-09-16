package database

import (
	"emailn/internal/domain/campaign"
)
 
type CampaignRepository struct{
    campaigns []campaign.Campaign
}

func (c *CampaignRepository) Save(campaign *campaign.Campaign) error {
    c.campaigns = append(c.campaigns, *campaign)
    return nil
}

func (c *CampaignRepository) Get() []campaign.Campaign {
    if c.campaigns == nil {
        return []campaign.Campaign{}
    }

    return c.campaigns
}
