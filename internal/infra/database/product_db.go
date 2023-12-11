package database

import (
	"github.com/rafaelsouzaribeiro/9-API/internal/entity"
	"gorm.io/gorm"
)

type Product struct {
	DB *gorm.DB
}

func NewProduct(db *gorm.DB) *Product {
	return &Product{DB: db}
}

func (p *Product) Create(product *entity.Product) error {
	return p.DB.Create(product).Error
}

func (p *Product) FindById(id string) (*entity.Product, error) {
	var product entity.Product
	erro := p.DB.First(&product, "id=?", id).Error

	return &product, erro

}

// Se o save estiver vazio ele salva mesmo assim
// precisa tratar para ver se o item existe
func (p *Product) Update(product *entity.Product) error {
	_, err := p.FindById(product.Id.String())

	if err != nil {
		return err
	}

	return p.DB.Save(product).Error
}

func (p *Product) Delete(id string) error {
	product, err := p.FindById(id)

	if err != nil {
		return err
	}

	return p.DB.Delete(product).Error
}

func (p *Product) FindAll(page, limit int, sort string) ([]entity.Product, error) {
	var products []entity.Product
	var err error

	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	if page != 0 && limit != 0 {
		err = p.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at " + sort).Find(&products).Error

	} else {
		err = p.DB.Order("created_at " + sort).Find(&products).Error
	}

	return products, err
}
