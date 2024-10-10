package entity

import "errors"

type Order struct {
	ID             int
	OrganizationID int
	DeliverID      int
	PaymentMethod  string
	Description    string
	Info           string
	Neighborhood   string
	Street         string
	Number         string
	Cep            string
	Paid           bool
}

func NewOrder(
	ID,
	OrganizationID,
	DeliverID int,
	PaymentMethod,
	Description,
	Info,
	Neighborhood,
	Street,
	Number,
	Cep string,
	Paid bool,
) (*Order, error) {
	order := Order{
		ID:             ID,
		OrganizationID: OrganizationID,
		DeliverID:      DeliverID,
		PaymentMethod:  PaymentMethod,
		Description:    Description,
		Info:           Info,
		Neighborhood:   Neighborhood,
		Street:         Street,
		Number:         Number,
		Cep:            Cep,
		Paid:           Paid,
	}
	if err := order.isValid(); err != nil {
		return nil, err
	}

	return &order, nil
}

func (o *Order) isValid() error {
	if o.OrganizationID == 0 {
		return errors.New("OrganizationID is zero")
	}
	if o.DeliverID == 0 {
		return errors.New("DeliverID is zero")
	}
	if o.PaymentMethod == "" {
		return errors.New("PaymentMethod is null")
	}
	if o.Description == "" {
		return errors.New("Description is null")
	}
	if o.Neighborhood == "" {
		return errors.New("Neighborhood is null")
	}
	if o.Street == "" {
		return errors.New("Street is null")
	}
	if o.Number == "" {
		return errors.New("Number is null")
	}
	if o.Cep == "" {
		return errors.New("Cep is null")
	}

	return nil
}
