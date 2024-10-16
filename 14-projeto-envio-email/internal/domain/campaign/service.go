package campaign

import (
	"emailn/internal/contract"
	internalerrors "emailn/internal/internal-errors"
	"encoding/json"
	"errors"

	"gorm.io/gorm"
)

type Service interface{
    Create(newCampaign contract.NewCampaign) (string, error)
    GetById(id string) (*contract.CampaignResponse, error)
    Delete(id string) error
}

type ServiceImp struct{
    Repository Repository
}
func (s *ServiceImp) Create(newCampaign contract.NewCampaign) (string, error) {
    campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails, newCampaign.CreatedBy)

    if err != nil {
        return "", err
    }

    err = s.Repository.Create(campaign)

    if err != nil {
        if !errors.Is(err, gorm.ErrRecordNotFound) {
            return "", internalerrors.ErrInternal
        }
        return "", err
    }

    return campaign.Id, err
}

func (s *ServiceImp) GetById(id string) (*contract.CampaignResponse, error) {
    campaign, err := s.Repository.GetById(id)
    if err != nil {
        if !errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, internalerrors.ErrInternal
        }
        return nil, err
    }

    return &contract.CampaignResponse{
        Id: campaign.Id,
        Status: campaign.Status,
        Content: campaign.Content,
        AmountOfEmailsToSend: len(campaign.Contacts),
        Name: campaign.Name,
        CreatedBy: campaign.CreatedBy,
    }, nil
}

func (s *ServiceImp) Delete(id string) error {
    campaign, err := s.GetById(id)

    if err != nil {
        return err
    }

    if campaign.Status != Pending {
        return errors.New("Campaign status Invalid")
    }

    var b Campaign
    js, _ := json.Marshal(campaign)
    json.Unmarshal(js, &b)
    b.Delete()

    err = s.Repository.Delete(&b)


    if err != nil {
        if !errors.Is(err, gorm.ErrRecordNotFound) {
            return internalerrors.ErrInternal
        }
        return err
    }

    return nil
}
