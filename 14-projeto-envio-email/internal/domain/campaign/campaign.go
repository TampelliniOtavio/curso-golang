package campaign

import (
	internalerrors "emailn/internal/internal-errors"
	"time"

	"github.com/rs/xid"
)

const (
    Pending string = "Pending"
    Started string = "Started"
    Done    string = "Done"
)

type Contact struct {
    Email string    `validate:"email"`
}

type Campaign struct {
    Id        string    `validate:"required"`
    Name      string    `validate:"min=5,max=24"`
    CreatedOn time.Time `validate:"required"`
    Status    string    `validate:"required"`
    Content   string    `validate:"min=5,max=1024"`
    Contacts  []Contact `validate:"min=1,dive"`
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {
    contacts := make([]Contact, len(emails))
    for i, email := range emails {
        contacts[i].Email = email
    }

    campaign := &Campaign{
        Id: xid.New().String(),
        Name: name,
        Content: content,
        CreatedOn: time.Now(),
        Status: Pending,
        Contacts: contacts,
    }

    err := internalerrors.ValidateStruct(campaign)
    if err != nil {
        return nil, err
    }

    return campaign, nil
}
