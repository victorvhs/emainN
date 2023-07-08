package campaign

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCampaign(t *testing.T) {
	assert := assert.New(t)

	name := "Minha campanha"
	content := "Conteúdo da minha campanha"
	shortDescription := "Descrição da minha campanha"
	contacts := []string{"email1@a.com", "email2@a.com", "email3@a.com"}

	campaign := NewCampaign(name, content, shortDescription, contacts)

	assert.Equal("1", campaign.ID)
	assert.Equal(name, campaign.Name)
	assert.Equal(content, campaign.Content)
	assert.Equal(shortDescription, campaign.ShortDescription)
	assert.Equal(len(contacts), len(campaign.Contacts))

}
