package campaign

import (
	"emailn/internal/contract"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

var (
	newCampaign = contract.NewCampaign{
		Name:             "Minha campanha",
		Content:          "Conteúdo da minha campanha",
		ShortDescription: "Descrição da minha campanha",
		Emails:           []string{"a@a.com", "b@b.com"},
	}
	mockR   = new(repositoryMock)
	service = Service{mockR}
)

func Test_Service_Create_Campaign(t *testing.T) {
	assert := assert.New(t)
	id, err := service.Create(newCampaign)
	assert.NotNil(id)
	assert.Nil(err)
}

func Test_Service_Validate_Domain_Error(t *testing.T) {
	assert := assert.New(t)
	newCampaign.Name = ""
	_, err := service.Create(newCampaign)

	assert.NotNil(err)
	assert.Equal("name should not be empty", err)
}

func Test_Service_Save_Campaign(t *testing.T) {
	mockR.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		return campaign.Name == newCampaign.Name &&
			campaign.Content == newCampaign.Content &&
			campaign.ShortDescription == newCampaign.ShortDescription &&
			len(campaign.Contacts) == len(newCampaign.Emails)

	})).Return(nil)

	service.Create(newCampaign)

	mockR.AssertExpectations(t)
}

func Test_Service_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(errors.New("error to save on database"))
	service.Repository = repositoryMock

	_, err := service.Create(newCampaign)

	assert.Equal("error to save on database", err.Error())

}
