package campaign

import (
	"emailn/internal/contract"
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
	newCamping = contract.NewCampaign{
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
	id, err := service.Create(newCamping)
	assert.NotNil(id)
	assert.Nil(err)
}
func Test_Service_Validate_Domain_Error(t *testing.T) {
	assert := assert.New(t)
	newCamping.Name = ""
	_, err := service.Create(newCamping)

	assert.NotNil(err)
	assert.Equal("name should not be empty", err)
}

func Test_Service_Save_Campaign(t *testing.T) {
	mockR.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		return campaign.Name == newCamping.Name &&
			campaign.Content == newCamping.Content &&
			campaign.ShortDescription == newCamping.ShortDescription &&
			len(campaign.Contacts) == len(newCamping.Emails)

	})).Return(nil)

	service.Create(newCamping)

	mockR.AssertExpectations(t)
}
