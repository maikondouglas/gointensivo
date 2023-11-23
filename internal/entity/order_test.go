package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfItGetsAnErrorIfIDIsBlank(t *testing.T) {
	order := Order{}

	assert.Error(t, order.Validate(), "id cannot be empty")
}

func TestIfItGetsAnErrorIfPriceIsBlank(t *testing.T) {
	order := Order{ID: "123"}

	assert.Error(t, order.Validate(), "price cannot be negative")
}

func TestIfItGetsAnErrorIfPriceIsZero(t *testing.T) {
	order := Order{ID: "123", Price: 0}

	assert.Error(t, order.Validate(), "price cannot be negative")
}

func TestIfItGetsAnErrorIfTaxLessThanZero(t *testing.T) {
	order := Order{ID: "123", Tax: -1}

	assert.Error(t, order.Validate(), "tax cannot be negative")
}

func TestFinalPrice(t *testing.T) {
	order := Order{ID: "123", Price: 10.0, Tax: 1.0}

	assert.NoError(t, order.Validate())
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 1.0, order.Tax)

	order.CalculateFinalPrice()

	assert.Equal(t, 11.0, order.FinalPrice)
}
