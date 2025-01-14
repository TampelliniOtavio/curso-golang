package campaign

import (
	internalerrors "emailn/internal/internal-errors"
	"time"

	"github.com/rs/xid"
)

const (
	Pending  string = "Pending"
	Started  string = "Started"
	Canceled string = "Canceled"
	Deleted  string = "Deleted"
	Failed   string = "Failed"
	Done     string = "Done"
)

type Contact struct {
	Id         string `validate:"required" gorm:"size:50"`
	Email      string `validate:"email" gorm:"size:100"`
	CampaignId string `validate:"required" gorm:"size:50"`
}

type Campaign struct {
	Id        string    `validate:"required" gorm:"size:50"`
	Name      string    `validate:"min=5,max=24" gorm:"size:100"`
	CreatedOn time.Time `validate:"required"`
	Status    string    `validate:"required" gorm:"size:20"`
	Content   string    `validate:"min=5,max=1024" gorm:"size:1024"`
	Contacts  []Contact `validate:"min=1,dive"`
	CreatedBy string    `validate:"required" gorm:"size:100"`
}

func (c *Campaign) Cancel() {
	c.Status = Canceled
}

func (c *Campaign) Delete() {
	c.Status = Deleted
}

func (c *Campaign) Fail() {
	c.Status = Failed
}

func (c *Campaign) Done() {
	c.Status = Done
}

func (c *Campaign) Started() {
	c.Status = Started
}

func NewCampaign(name string, content string, emails []string, createdBy string) (*Campaign, error) {
	campaignId := xid.New().String()
	contacts := make([]Contact, len(emails))
	for i, email := range emails {
		contacts[i].Email = email
		contacts[i].Id = xid.New().String()
		contacts[i].CampaignId = campaignId
	}

	campaign := &Campaign{
		Id:        campaignId,
		Name:      name,
		Content:   content,
		CreatedOn: time.Now(),
		Status:    Pending,
		Contacts:  contacts,
		CreatedBy: createdBy,
	}

	err := internalerrors.ValidateStruct(campaign)
	if err != nil {
		return nil, err
	}

	return campaign, nil
}
