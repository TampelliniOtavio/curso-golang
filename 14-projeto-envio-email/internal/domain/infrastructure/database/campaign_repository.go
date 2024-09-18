package database

import (
	"emailn/internal/domain/campaign"
	"errors"
)
 
type CampaignRepository struct{
    campaigns []campaign.Campaign
}

func (c *CampaignRepository) Save(campaign *campaign.Campaign) error {
    c.campaigns = append(c.campaigns, *campaign)
    return nil
}

func (c *CampaignRepository) Get() ([]campaign.Campaign, error) {
    if c.campaigns == nil {
        return []campaign.Campaign{}, nil
    }

    return c.campaigns, nil
}

func (c *CampaignRepository) GetById(id string) (*campaign.Campaign, error) {
    for _, v := range c.campaigns {
        if v.Id == id {
            return &v, nil
        }
    }

    return nil, errors.New("Not Found")
}
