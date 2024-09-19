package mock

import (
	"emailn/internal/contract"

	"github.com/stretchr/testify/mock"
)

type CampaignServiceMock struct{
    mock.Mock
}
func (m *CampaignServiceMock) Create(newCampaign contract.NewCampaign) (string, error) {
    args := m.Called(newCampaign)
    return args.String(0), args.Error(1)
}

func (m *CampaignServiceMock) GetById(id string) (*contract.CampaignResponse, error) {
    args := m.Called(id)
    err := args.Error(1)

    if err != nil {
        return nil, err
    }

    return args.Get(0).(*contract.CampaignResponse), nil
}

func (m *CampaignServiceMock) Cancel(id string) error {
    args := m.Called(id)
    return args.Error(0)
}
