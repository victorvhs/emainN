package campaign

import (
	"github.com/rs/xid"
	"time"
)

type Contact struct {
	Email string
}
type Campaign struct {
	ID               string
	Name             string
	CreatedOn        time.Time
	ShortDescription string
	Content          string
	Contacts         []Contact
}

func NewCampaign(name string, content string, shortDescription string, emails []string) *Campaign {
	contacts := make([]Contact, len(emails))
	for i, email := range emails {
		contacts[i].Email = email
	}
	return &Campaign{
		ID:               xid.New().String(),
		Name:             name,
		Content:          content,
		ShortDescription: shortDescription,
		CreatedOn:        time.Now(),
		Contacts:         contacts,
	}
}
