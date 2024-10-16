package usecase

import "github.com/ThalesMonteir0/entregue-ja/internal/entity"

type OrganizationDTO struct {
	FantasyName  string `json:"fantasy_name"`
	SocialReason string `json:"social_reason"`
	Cnpj         string `json:"cnpj"`
	ContactOwner string `json:"contact_owner"`
	EmailOwner   string `json:"email_owner"`
}

type OrganizationDTOResponse struct {
	ID           int    `json:"id"`
	FantasyName  string `json:"fantasy_name"`
	SocialReason string `json:"social_reason"`
	Cnpj         string `json:"cnpj"`
	ContactOwner string `json:"contact_owner"`
	EmailOwner   string `json:"email_owner"`
}

type CreateOrganizationUseCase struct {
	Repository entity.OrganizationRepositoryInterface
}

func NewCreateOrganizationUseCase(repository entity.OrganizationRepositoryInterface) CreateOrganizationUseCase {
	return CreateOrganizationUseCase{
		repository,
	}
}

func (c CreateOrganizationUseCase) Execute(organizationDTO OrganizationDTO) (OrganizationDTOResponse, error) {
	organization, err := entity.NewOrganization(
		organizationDTO.FantasyName,
		organizationDTO.SocialReason,
		organizationDTO.Cnpj,
		organizationDTO.ContactOwner,
		organizationDTO.EmailOwner,
	)
	if err != nil {

	}

	if err = c.Repository.Save(organization); err != nil {
		return OrganizationDTOResponse{}, err
	}

	organizationDTOResponse := OrganizationDTOResponse{
		ID:           organization.ID,
		FantasyName:  organization.FantasyName,
		SocialReason: organization.SocialReason,
		Cnpj:         organization.Cnpj,
		ContactOwner: organization.ContactOwner,
		EmailOwner:   organization.EmailOwner,
	}

	return organizationDTOResponse, nil
}
