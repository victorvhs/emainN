package campaign

import "testing"

func TestNewCampaign(t *testing.T) {
	name := "Minha campanha"
	content := "Conteúdo da minha campanha"
	shortDescription := "Descrição da minha campanha"
	contacts := []string{"email1@a.com", "email2@a.com", "email3@a.com"}
	campaign := NewCampaign(name, content, shortDescription, contacts)

	if campaign.Name != name {
		t.Errorf("Expected campaign name to be %s, but got %s", name, campaign.Name)
	} else if campaign.ID != "1" {
		t.Errorf("Expected campaign ID to be 1, but got %s", campaign.ID)
	} else if campaign.Content != content {
		t.Errorf("Expected campaign content to be %s, but got %s", content, campaign.Content)
	} else if len(campaign.Contacts) != len(contacts) {
		t.Errorf("Expected campaign contacts %d have, but got %d", len(contacts), len(campaign.Contacts))
	} else if campaign.ShortDescription != shortDescription {
		t.Errorf("Expected campaign shortDescription to be %s, but got %s", shortDescription, campaign.ShortDescription)
	}

}
