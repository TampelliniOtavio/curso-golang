package campaign

import (
	internalerrors "emailn/internal/internal-errors"
	"encoding/json"
	"errors"

	"gorm.io/gorm"
)

type Service interface {
	Create(newCampaign NewCampaignRequest) (string, error)
	GetById(id string) (*CampaignResponse, error)
	Delete(id string) error
	Start(id string) error
}

type ServiceImp struct {
	Repository Repository
	SendMail   func(Campaign *Campaign) error
}

func (s *ServiceImp) Create(newCampaign NewCampaignRequest) (string, error) {
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

func (s *ServiceImp) GetById(id string) (*CampaignResponse, error) {
	campaign, err := s.Repository.GetById(id)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, internalerrors.ErrInternal
		}
		return nil, err
	}

	return &CampaignResponse{
		Id:                   campaign.Id,
		Status:               campaign.Status,
		Content:              campaign.Content,
		AmountOfEmailsToSend: len(campaign.Contacts),
		Name:                 campaign.Name,
		CreatedBy:            campaign.CreatedBy,
	}, nil
}

func (s *ServiceImp) getAndValidateStatusIsPending(id string) (*Campaign, error) {
	campaign, err := s.Repository.GetById(id)

	if err != nil {
		return nil, err
	}

	if campaign.Status != Pending {
		return nil, errors.New("Campaign status Invalid")
	}

	return campaign, nil
}

func (s *ServiceImp) Delete(id string) error {
	campaign, err := s.getAndValidateStatusIsPending(id)

	if err != nil {
		return err
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

func (s *ServiceImp) SendEmailAndUpdateStatus(campaignSaved *Campaign) {
	err := s.SendMail(campaignSaved)
	if err != nil {
		campaignSaved.Fail()
	} else {
		campaignSaved.Done()
	}

	s.Repository.Update(campaignSaved)
}

func (s *ServiceImp) Start(id string) error {
	campaignSaved, err := s.getAndValidateStatusIsPending(id)

	if err != nil {
		return err
	}

	campaignSaved.Started()
	err = s.Repository.Update(campaignSaved)

	if err != nil {
		return internalerrors.ErrInternal
	}

	return nil
}
