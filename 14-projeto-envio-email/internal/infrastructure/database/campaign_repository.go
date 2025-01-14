package database

import (
	"emailn/internal/domain/campaign"

	"gorm.io/gorm"
)

type CampaignRepository struct {
	Db *gorm.DB
}

func (c *CampaignRepository) Create(campaign *campaign.Campaign) error {
	tx := c.Db.Create(campaign)
	return tx.Error
}

func (c *CampaignRepository) Update(campaign *campaign.Campaign) error {
	tx := c.Db.Save(campaign)
	return tx.Error
}

func (c *CampaignRepository) Get() (*[]campaign.Campaign, error) {
	var campaigns []campaign.Campaign

	tx := c.Db.Find(&campaigns)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &campaigns, nil
}

func (c *CampaignRepository) GetById(id string) (*campaign.Campaign, error) {
	var campaign campaign.Campaign

	tx := c.Db.Preload("Contacts").First(&campaign, "id = ?", id)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &campaign, nil
}

func (c *CampaignRepository) Delete(campaign *campaign.Campaign) error {
	tx := c.Db.Select("Contacts").Delete(&campaign)

	return tx.Error
}

func (c *CampaignRepository) GetCampaignsToBeSent() ([]campaign.Campaign, error) {
	var campaigns []campaign.Campaign
	tx := c.Db.Preload("Contacts").Find(
		&campaigns,
		"status = ? and date_part('minute', now()::timestamp - updated_on::timestamp) >= ?", // diferença em minutos
		campaign.Started,
		1,
	)
	return campaigns, tx.Error
}
