package entity

import "errors"

type Order struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

func NewOrder(id string, price float64, tax float64) (*Order, error) {
	order := &Order{
		ID:    id,
		Price: price,
		Tax:   tax,
	}

	if err := order.Validate(); err != nil {
		return nil, err
	}

	return order, nil
}

func (o *Order) Validate() error {
	if o.ID == "" {
		return errors.New("id cannot be empty")
	}
	if o.Price <= 0 {
		return errors.New("price cannot be negative")
	}
	if o.Tax < 0 {
		return errors.New("tax cannot be negative")
	}

	return nil
}

func (o *Order) CalculateFinalPrice() error {
	o.FinalPrice = o.Price + o.Tax

	return o.Validate()
}
