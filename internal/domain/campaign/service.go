package campaign

import "emailn/internal/contract"

type Service struct {
	Repository Repository
}

func (s *Service) Create(contract.NewCampaign) error {

}
