package entity

import (
	"errors"
	"time"

	"github.com/rafaelsouzaribeiro/9-API/pkg/entity"
)

var (
	ErrIdIsRequired    = errors.New("Id obrigatório")
	ErrNameIsRequired  = errors.New("Nome obrigatório")
	ErrPriceIsRequired = errors.New("Preço obrigatório")
	ErrInvalidId       = errors.New("Invalido Id")
	ErrInvalidPrice    = errors.New("Invalido Preço")
)

type Product struct {
	Id        entity.Id `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price float64) (*Product, error) {
	product := &Product{
		Id:        entity.NewId(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}

	err := product.Validate()

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *Product) Validate() error {
	if p.Id.String() == "" {
		return ErrIdIsRequired
	}

	if _, err := entity.ParseId(p.Id.String()); err != nil {
		return ErrInvalidId
	}

	if p.Name == "" {
		return ErrNameIsRequired
	}

	if p.Price == 0 {
		return ErrInvalidPrice
	}

	if p.Price < 0 {
		return ErrInvalidPrice
	}

	return nil
}
