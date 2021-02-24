package campaign

type Service interface {
	GetCampaigns(userID int) ([]Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCampaigns(userID int) ([]Campaign, error) {
	if userID != 0 {
		userCampaigns, err := s.repository.FindByUserID(userID)
		if err != nil {
			return userCampaigns, err
		}

		return userCampaigns, nil
	}

	campaigns, err := s.repository.FindAll()
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}
