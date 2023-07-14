package campaign

import (
	"emailn/internal/contract"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Service_Create_Campaign(t *testing.T) {
	assert := assert.New(t)
	service := Service{}
	newCamping := contract.NewCampaign{
		Name:             "Minha campanha",
		Content:          "Conteúdo da minha campanha",
		ShortDescription: "Descrição da minha campanha",
		Emails:           []string{"a@a.com", "b@b.com"},
	}
	id, err := service.Create(newCamping)

	assert.NotNil(id)
	assert.Nil(err)
}

func Test_Service_Save_Campaign(t *testing.T) {
	assert := assert.New(t)
	service := Service{}
	newCamping := contract.NewCampaign{
		Name:             "Minha campanha",
		Content:          "Conteúdo da minha campanha",
		ShortDescription: "Descrição da minha campanha",
		Emails:           []string{"a@a.com", "b@b.com"},
	}
	id, err := service.Create(newCamping)

	assert.NotNil(id)
	assert.Nil(err)
}
