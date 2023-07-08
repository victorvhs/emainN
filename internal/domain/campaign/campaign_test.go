package campaign

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_NewCampaign_Create_Campaign(t *testing.T) {
	assert := assert.New(t)

	name := "Minha campanha"
	content := "Conteúdo da minha campanha"
	shortDescription := "Descrição da minha campanha"
	contacts := []string{"email1@a.com", "email2@a.com", "email3@a.com"}

	campaign := NewCampaign(name, content, shortDescription, contacts)

	assert.Equal(name, campaign.Name)
	assert.Equal(content, campaign.Content)
	assert.Equal(shortDescription, campaign.ShortDescription)
	assert.Equal(len(contacts), len(campaign.Contacts))

}

func Test_NewCampaing_Not_Nil_ID(t *testing.T) {
	assert := assert.New(t)

	name := "Minha campanha"
	content := "Conteúdo da minha campanha"
	shortDescription := "Descrição da minha campanha"
	contacts := []string{"email1@a.com", "email2@a.com", "email3@a.com"}

	campaign := NewCampaign(name, content, shortDescription, contacts)

	assert.NotEmpty(campaign.ID, "ID should not be empty")
	assert.NotNilf(campaign.ID, "ID should not be nil")

}

func Test_NewCampaing_Not_Nil_CreatedOn(t *testing.T) {
	assert := assert.New(t)

	name := "Minha campanha"
	content := "Conteúdo da minha campanha"
	shortDescription := "Descrição da minha campanha"
	contacts := []string{"email1@a.com", "email2@a.com", "email3@a.com"}
	now := time.Now().Add(-time.Minute)

	campaign := NewCampaign(name, content, shortDescription, contacts)

	assert.NotEmpty(campaign.CreatedOn, "CreatedOn should not be empty")
	assert.NotNilf(campaign.CreatedOn, "CreatedOn should not be nil")
	assert.GreaterOrEqual(campaign.CreatedOn, now, "CreatedOn should be greater or equal than now")

}
