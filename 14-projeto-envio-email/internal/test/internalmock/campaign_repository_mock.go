package internalmock

import (
	"emailn/internal/domain/campaign"

	"github.com/stretchr/testify/mock"
)

type CampaignRepositoryMock struct{
    mock.Mock
}

func (c *CampaignRepositoryMock) Create(campaign *campaign.Campaign) error {
    args := c.Called(campaign)
    return args.Error(0)
}

func (c *CampaignRepositoryMock) Update(campaign *campaign.Campaign) error {
    args := c.Called(campaign)
    return args.Error(0)
}

func (c *CampaignRepositoryMock) Get() (*[]campaign.Campaign, error) {
    args := c.Called()
    first := args.Get(0)
    err := args.Error(1)

    if first == nil {
        return nil, err
    }
    return args.Get(0).(*[]campaign.Campaign), args.Error(1)
}

func (c *CampaignRepositoryMock) GetById(id string) (*campaign.Campaign, error) {
    args := c.Called(id)
    first := args.Get(0)
    err := args.Error(1)

    if first == nil {
        return nil, err
    }
    return args.Get(0).(*campaign.Campaign), args.Error(1)
}

func (c *CampaignRepositoryMock) Delete(campaign *campaign.Campaign) error {
    args := c.Called(campaign)
    return args.Error(0)
}
