package entity

type OrganizationRepositoryInterface interface {
	Save(organizationDTO *Organization) error
}
