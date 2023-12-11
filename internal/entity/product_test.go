package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	p, err := NewProduct("Geladeira", 10)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.Id)
	assert.Equal(t, "Geladeira", p.Name)
	assert.Equal(t, 10.0, p.Price)

}

func TestProductWhenNameIsRequired(t *testing.T) {
	p, err := NewProduct("", 10)
	assert.Nil(t, p)
	assert.Equal(t, ErrNameIsRequired, err)

}

func TestProductWhenPriceIsRequired(t *testing.T) {
	p, err := NewProduct("Geladeira", 0)
	assert.Nil(t, p)
	assert.Equal(t, ErrInvalidPrice, err)

}

func TestProductWhenPriceIsValid(t *testing.T) {
	p, err := NewProduct("Geladeira", -10)
	assert.Nil(t, p)
	assert.Equal(t, ErrInvalidPrice, err)

}

func TestProductValidate(t *testing.T) {
	p, err := NewProduct("Geladeira", 10)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Nil(t, p.Validate())

}
