package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/rafaelsouzaribeiro/9-API/internal/dto"
	"github.com/rafaelsouzaribeiro/9-API/internal/entity"
	"github.com/rafaelsouzaribeiro/9-API/internal/infra/database"
	pkg "github.com/rafaelsouzaribeiro/9-API/pkg/entity"
)

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

// Create Product godoc
// @Summary      Create product
// @Description  Create products
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        request     body      dto.CreateProductInput  true  "product request"
// @Success      201
// @Failure      500         {object}  Error
// @Router       /products [post]
// @Security ApiKeyAuth
func (p *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	// if r.Method != http.MethodPost {
	// 	http.Error(w, "Method not allow", http.StatusMethodNotAllowed)
	// 	return
	// }

	var product dto.CreateProductInput

	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	ps, errs := entity.NewProduct(product.Name, product.Price)

	if errs != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	err = p.ProductDB.Create(ps)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	message := []byte("Criado com sucesso!\n")
	w.Write(message)
}

// GetProduct godoc
// @Summary      Get a product
// @Description  Get a product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "product ID" Format(uuid)
// @Success      200  {object}  entity.Product
// @Failure      404
// @Failure      500  {object}  Error
// @Router       /products/{id} [get]
// @Security ApiKeyAuth
func (p *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	product, err := p.ProductDB.FindById(id)

	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(product)
}

// UpdateProduct godoc
// @Summary      Update a product
// @Description  Update a product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id        	path      string                  true  "product ID" Format(uuid)
// @Param        request    body      dto.CreateProductInput  true  "product request"
// @Success      200
// @Failure      404
// @Failure      500       {object}  Error
// @Router       /products/{id} [put]
// @Security ApiKeyAuth
func (p *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	product.Id, err = pkg.ParseId(id)

	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	_, err = p.ProductDB.FindById(id)

	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	err = p.ProductDB.Update(&product)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	message := []byte("Utualizado com sucesso!\n")
	w.Write(message)

}

// DeleteProduct godoc
// @Summary      Delete a product
// @Description  Delete a product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id        path     string   true  "product ID" Format(uuid)
// @Success      200
// @Failure      404
// @Failure      500       {object}  Error
// @Router       /products/{id} [delete]
// @Security ApiKeyAuth
func (u *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	_, err := u.ProductDB.FindById(id)

	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	err = u.ProductDB.Delete(id)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	message := []byte("Deletado com sucesso!\n")
	w.Write(message)
}

// ListAccounts godoc
// @Summary      List products
// @Description  get all products
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        page      query     string  false  "page number"
// @Param        limit     query     string  false  "limit"
// @Success      200       {array}   entity.Product
// @Failure      404       {object}  Error
// @Failure      500       {object}  Error
// @Router       /products [get]
// @Security ApiKeyAuth
func (u *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")

	pageInt, err := strconv.Atoi(page)

	if err != nil {
		pageInt = 0
	}

	limitInt, err := strconv.Atoi(limit)

	if err != nil {
		limitInt = 0
	}

	sort := r.URL.Query().Get("sort")
	products, errs := u.ProductDB.FindAll(pageInt, limitInt, sort)

	if errs != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(products)

}
