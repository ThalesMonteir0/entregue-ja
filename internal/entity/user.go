package entity

import "errors"

type User struct {
	ID             int
	OrganizationID int
	TypeUserID     int
	Name           string
	Password       string
	Role           string
	Email          string
}

func NewUser(organizationID, typeUserID int, name string, password, role, email string) (*User, error) {
	user := User{
		OrganizationID: organizationID,
		Name:           name,
		Password:       password,
		Role:           role,
		TypeUserID:     typeUserID,
		Email:          email,
	}

	if err := user.isValid(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *User) isValid() error {
	if u.Name == "" {
		return errors.New("Name is required")
	}
	if u.OrganizationID == 0 {
		return errors.New("OrganizationID is required")
	}
	if u.Role == "" {
		return errors.New("Role is required")
	}
	if u.Password == "" {
		return errors.New("Password is required")
	}

	return nil
}

func (u *User) EncryptPassword() error {

}
