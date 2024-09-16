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

