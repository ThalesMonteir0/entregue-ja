package entity

import "errors"

type Organization struct {
	ID           int
	FantasyName  string
	SocialReason string
	Cnpj         string
	ContactOwner string
	EmailOwner   string
}

func NewOrganization(id int, fantasyName, socialReason, cnpj, contactOwner, emailOwner string) (*Organization, error) {
	org := Organization{
		ID:           id,
		FantasyName:  fantasyName,
		SocialReason: socialReason,
		Cnpj:         cnpj,
		ContactOwner: contactOwner,
		EmailOwner:   emailOwner,
	}

	if err := org.isValid(); err != nil {
		return nil, err
	}

	return &org, nil
}

func (o *Organization) isValid() error {
	if o.FantasyName == "" {
		return errors.New("organization.FantasyName is required")
	}
	if o.SocialReason == "" {
		return errors.New("organization.SocialReason is required")
	}
	if o.Cnpj == "" {
		return errors.New("organization.Cnpj is required")
	}
	if o.ContactOwner == "" {
		return errors.New("organization.ContactOwner is required")
	}
	if o.EmailOwner == "" {
		return errors.New("organization.EmailOwner is required")
	}

	return nil
}
