package internalmock

import (
	"emailn/internal/domain/campaign"

	"github.com/stretchr/testify/mock"
)

type CampaignServiceMock struct {
	mock.Mock
}

func (m *CampaignServiceMock) Create(newCampaign campaign.NewCampaignRequest) (string, error) {
	args := m.Called(newCampaign)
	return args.String(0), args.Error(1)
}

func (m *CampaignServiceMock) GetById(id string) (*campaign.CampaignResponse, error) {
	args := m.Called(id)
	err := args.Error(1)

	if err != nil {
		return nil, err
	}

	return args.Get(0).(*campaign.CampaignResponse), nil
}

func (m *CampaignServiceMock) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *CampaignServiceMock) Start(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
