package controller

import (
	"ProjectPhonePay/constants"
	"ProjectPhonePay/model"
	"ProjectPhonePay/service"
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type SellerController struct {
	SellerService service.ISellerService
}

func NewSellerController(seller service.ISellerService) SellerController {
	return SellerController{
		SellerService: seller,
	}
}

func (c *SellerController) SellerRegister(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newSeller model.SellerRegisterReq

	err := json.NewDecoder(r.Body).Decode(&newSeller)

	if err != nil {
		w.WriteHeader(400)
		_ = json.NewEncoder(w).Encode(err.Error())
		return
	}

	validate := validator.New()
	err = validate.Struct(newSeller)

	if err != nil {
		w.WriteHeader(400)
		_ = json.NewEncoder(w).Encode(err.Error())
		return
	}

	err = c.SellerService.SellerRegistration(newSeller)

	if err != nil {
		w.WriteHeader(400)
		_ = json.NewEncoder(w).Encode(err.Error())

		return
	}
	w.WriteHeader(200)
	_ = json.NewEncoder(w).Encode("Seller Register Successfully")
	return
}

func (c *SellerController) SellerLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var login model.SellerLoginReq
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		w.WriteHeader(400)
		_ = json.NewEncoder(w).Encode(err.Error())
		return
	}

	validate := validator.New()
	err = validate.Struct(login)

	if err != nil {
		w.WriteHeader(400)
		_ = json.NewEncoder(w).Encode(err.Error())
		return
	}

	err = c.SellerService.SellerLogin(login)

	if err != nil {
		w.WriteHeader(400)
		_ = json.NewEncoder(w).Encode(err.Error())

		return
	}
	w.WriteHeader(200)
	_ = json.NewEncoder(w).Encode("Seller Successfully LoggedIn")

	return
}

func (c *SellerController) SellerLogout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var logout model.SellerLogoutReq
	err := json.NewDecoder(r.Body).Decode(&logout)
	if err != nil {
		w.WriteHeader(400)
		_ = json.NewEncoder(w).Encode(err.Error())
		return
	}

	validate := validator.New()
	err = validate.Struct(logout)

	if err != nil {
		w.WriteHeader(400)
		_ = json.NewEncoder(w).Encode(err.Error())
		return
	}

	err = c.SellerService.SellerLogout(logout)

	if err != nil {
		w.WriteHeader(400)
		_ = json.NewEncoder(w).Encode(err.Error())

		return
	}

	_ = json.NewEncoder(w).Encode("Seller Successfully LoggedOut")
	w.WriteHeader(200)
	return
}

func (c *SellerController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product model.ProductCreate
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(400)
		_ = json.NewEncoder(w).Encode(err.Error())
		return
	}

	validate := validator.New()
	err = validate.Struct(product)

	if err != nil {
		w.WriteHeader(400)
		_ = json.NewEncoder(w).Encode(err.Error())
		return
	}

	if product.Type != constants.CategoryA && product.Type != constants.CategoryB && product.Type != constants.CategoryC {
		err = errors.New("type Should be in : Wearable , Grocery, Furniture")
		w.WriteHeader(400)
		_ = json.NewEncoder(w).Encode(err.Error())
		return
	}
	err = c.SellerService.CreateProduct(product.Seller.Email, product)

	if err != nil {
		w.WriteHeader(400)
		_ = json.NewEncoder(w).Encode(err.Error())

		return
	}

	w.WriteHeader(200)
	_ = json.NewEncoder(w).Encode("ProductCreate Created Successfully")

	return
}

func (c *SellerController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product model.ProductUpdate
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(400)
		_ = json.NewEncoder(w).Encode(err.Error())
		return
	}

	validate := validator.New()
	err = validate.Struct(product)

	if err != nil {
		w.WriteHeader(400)
		_ = json.NewEncoder(w).Encode(err.Error())
		return
	}

	if product.Type != "" && product.Type != constants.CategoryA && product.Type != constants.CategoryB && product.Type != constants.CategoryC {
		err = errors.New("type Should be in : Wearable , Grocery, Furniture")
		w.WriteHeader(400)
		_ = json.NewEncoder(w).Encode(err.Error())
		return
	}
	err = c.SellerService.UpdateProduct(product.Seller.Email, product)

	if err != nil {
		w.WriteHeader(400)
		_ = json.NewEncoder(w).Encode(err.Error())

		return
	}

	w.WriteHeader(200)
	_ = json.NewEncoder(w).Encode("ProductCreate Created Successfully")

	return
}

func (c *SellerController) GetAllSeller(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	seller := c.SellerService.GetAllSeller()
	b, _ := json.Marshal(seller)
	w.WriteHeader(200)
	_, _ = w.Write(b)
	_ = json.NewEncoder(w).Encode("All Seller Printed")

	return
}

func (c *SellerController) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	seller := c.SellerService.GetAllSeller()
	b, _ := json.Marshal(seller)
	w.WriteHeader(200)
	_, _ = w.Write(b)
	_ = json.NewEncoder(w).Encode("All Seller Printed")

	return
}
