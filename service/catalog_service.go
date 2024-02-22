package service

import "ProjectPhonePay/dtos"

type ICatalogService interface {
	CreateProduct(product *dtos.Product) (int64, error)
	UpdateProduct(product *dtos.Product) error
	AddProductQty(product *dtos.Product, qty int64) error
	DecreaseProductQty(product *dtos.Product, qty int64) error
	BuyProduct(productId int64, qty int64, email string) (int64, error)
	UpdateStatus(orderId int64, status string, email string) error
	GetProductById(ProdId int64) (*dtos.Product, error)
}
