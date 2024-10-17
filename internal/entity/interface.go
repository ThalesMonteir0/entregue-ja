package entity

type OrganizationRepositoryInterface interface {
	Save(organization *Organization) error
}

type UserRepositoryInterface interface {
	Save(user *User) error
}
