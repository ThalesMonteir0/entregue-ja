package entity

import "errors"

type Deliver struct {
	ID             int
	OrganizationID int
	Name           string
	Tel            string
}

func NewDeliverer(id, organizationID int, name string, tel string) (*Deliver, error) {
	deliverer := Deliver{
		ID:             id,
		OrganizationID: organizationID,
		Name:           name,
		Tel:            tel,
	}

	if err := deliverer.isValid(); err != nil {
		return nil, err
	}

	return &deliverer, nil
}

func (d *Deliver) isValid() error {
	if d.OrganizationID == 0 {
		return errors.New("OrganizationID is zero")
	}
	if d.Tel == "" {
		return errors.New("Tel is null")
	}
	if d.Name == "" {
		return errors.New("Name is null")
	}

	return nil
}
