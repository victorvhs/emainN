package campaign

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	name             = "Minha campanha"
	content          = "Conteúdo da minha campanha"
	shortDescription = "Descrição da minha campanha"
	contacts         = []string{"email1@a.com", "email2@a.com", "email3@a.com"}
)

func Test_NewCampaign_Create_Campaign(t *testing.T) {
	assertions := assert.New(t)
	campaign, _ := NewCampaign(name, content, shortDescription, contacts)

	assertions.Equal(name, campaign.Name)
	assertions.Equal(content, campaign.Content)
	assertions.Equal(shortDescription, campaign.ShortDescription)
	assertions.Equal(len(contacts), len(campaign.Contacts))

}

func Test_NewCamping_Not_Nil_ID(t *testing.T) {
	assertions := assert.New(t)
	campaign, _ := NewCampaign(name, content, shortDescription, contacts)

	assertions.NotEmpty(campaign.ID, "ID should not be empty")
	assertions.NotNilf(campaign.ID, "ID should not be nil")
}

func Test_NewCamping_Not_Nil_CreatedOn(t *testing.T) {
	assertions := assert.New(t)
	now := time.Now().Add(-time.Minute)

	campaign, _ := NewCampaign(name, content, shortDescription, contacts)

	assertions.NotEmpty(campaign.CreatedOn, "CreatedOn should not be empty")
	assertions.NotNilf(campaign.CreatedOn, "CreatedOn should not be nil")
	assertions.GreaterOrEqual(campaign.CreatedOn, now, "CreatedOn should be greater or equal than now")

}

func Test_NewCamping_Not_Nil_Name(t *testing.T) {
	assertions := assert.New(t)

	_, err := NewCampaign("", content, shortDescription, contacts)

	assertions.Equal("Name should not be empty", err.Error())

}

func Test_NewCamping_Not_Nil_Content(t *testing.T) {
	assertions := assert.New(t)

	_, err := NewCampaign(name, "", shortDescription, contacts)

	assertions.Equal("Content should not be empty", err.Error())

}

func Test_NewCamping_Not_Nil_Contacts(t *testing.T) {
	assertions := assert.New(t)

	_, err := NewCampaign(name, content, shortDescription, []string{})

	assertions.Equal("Contacts should not be empty", err.Error())

}
