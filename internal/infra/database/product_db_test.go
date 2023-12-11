package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/rafaelsouzaribeiro/9-API/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDatabaseProduct() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&entity.Product{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func TestCreateNewProduct(t *testing.T) {
	db, err := setupTestDatabaseProduct()

	if err != nil {
		t.Error(err)
	}

	product, err := entity.NewProduct("Product 1", 10.0)
	assert.NoError(t, err)
	products := NewProduct(db)
	err = products.Create(product)
	assert.NoError(t, err)
	assert.NotEmpty(t, product.Id)
}

func TestFindAllProduct(t *testing.T) {
	db, err := setupTestDatabaseProduct()

	if err != nil {
		t.Error(err)
	}

	for i := 1; i < 24; i++ {
		// fmt.Sprintf("Product %d",i)
		// Product 1... até 24
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		assert.NoError(t, err)
		db.Create(product)
	}

	productsDb := NewProduct(db)
	products, err := productsDb.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products, err = productsDb.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)

	products, err = productsDb.FindAll(3, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 3)
	assert.Equal(t, "Product 21", products[0].Name)
	assert.Equal(t, "Product 23", products[2].Name)

}

func TestFindById(t *testing.T) {
	db, err := setupTestDatabaseProduct()

	if err != nil {
		t.Error(err)
	}

	product, err := entity.NewProduct("Product 1", 10.0)
	assert.NoError(t, err)
	db.Create(product)
	products := NewProduct(db)
	product, err = products.FindById(product.Id.String())
	assert.NoError(t, err)
	assert.Equal(t, "Product 1", product.Name)

}

func TestUpdateProduct(t *testing.T) {
	db, err := setupTestDatabaseProduct()

	if err != nil {
		t.Error(err)
	}

	product, err := entity.NewProduct("Product 1", 10.0)
	assert.NoError(t, err)
	db.Create(product)
	productDb := NewProduct(db)
	product.Name = "Product 2"
	err = productDb.Update(product)
	assert.NoError(t, err)
	product, err = productDb.FindById(product.Id.String())
	assert.NoError(t, err)
	assert.Equal(t, "Product 2", product.Name)
}

func TestDelete(t *testing.T) {
	db, err := setupTestDatabaseProduct()

	if err != nil {
		t.Error(err)
	}

	product, err := entity.NewProduct("Product 1", 10.0)
	assert.NoError(t, err)
	db.Create(product)

	productDb := NewProduct(db)
	err = productDb.Delete(product.Id.String())
	assert.NoError(t, err)

	_, errs := productDb.FindById(product.Id.String())
	assert.Error(t, errs)

}
