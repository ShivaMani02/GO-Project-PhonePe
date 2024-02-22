package service

import (
	"ProjectPhonePay/dtos"
	"ProjectPhonePay/model"
)

type ISellerService interface {
	SellerLogin(login model.SellerLoginReq) error
	SellerLogout(login model.SellerLogoutReq) error
	SellerRegistration(user model.SellerRegisterReq) error
	GetAllSeller() []dtos.Seller
	UpdateStatus(email string, newStatus string, orderId int64) error
	CreateProduct(email string, product model.ProductCreate) error
	UpdateProduct(email string, product model.ProductUpdate) error
	AddProductQty(email string, product model.ProductUpdate) error
}
