package service

import (
	"ProjectPhonePay/dtos"
	"ProjectPhonePay/model"
)

type IUserService interface {
	UserLogin(login model.UserLoginReq) error
	UserLogout(login model.UserLogoutReq) error
	UserRegistration(user model.UserRegisterReq) error
	GetProducts(email string, filter string) (model.GetProductsResponse, error)
	GetOrders(email string) (model.GetOrderResponse, error)
	AddToCart(email string, prodId int64, qty int64) error
	GetAllUser() []dtos.User
	Checkout(email string) (string, string, error)
}
