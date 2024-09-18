package campaign

import (
	"emailn/internal/contract"
	internalerrors "emailn/internal/internal-errors"
)
type Service interface{
    Create(newCampaign contract.NewCampaign) (string, error)
    GetById(id string) (*contract.CampaignResponse, error)
}

type ServiceImp struct{
    Repository Repository
}
func (s *ServiceImp) Create(newCampaign contract.NewCampaign) (string, error) {
    campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)

    if err != nil {
        return "", err
    }

    err = s.Repository.Save(campaign)

    if err != nil {
        return "", internalerrors.ErrInternal
    }

    return campaign.Id, err
}

func (s *ServiceImp) GetById(id string) (*contract.CampaignResponse, error) {
    campaign, err := s.Repository.GetById(id)
    if err != nil {
        return nil, internalerrors.ErrInternal
    }

    return &contract.CampaignResponse{
        Id: campaign.Id,
        Status: campaign.Status,
        Content: campaign.Content,
        Name: campaign.Name,
    }, nil
}
