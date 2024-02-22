package service_impl

import (
	"ProjectPhonePay/constants"
	"ProjectPhonePay/dtos"
	"ProjectPhonePay/model"
	"errors"
	"fmt"
	"strconv"
)

type CatalogServiceImpl struct {
	Products map[int64]dtos.Product
	Orders   map[int64]dtos.Order
}

var productId int64
var orderId int64

func NewCatalogService() *CatalogServiceImpl {
	productId = 0
	orderId = 0

	return &CatalogServiceImpl{
		Products: make(map[int64]dtos.Product, 0),
		Orders:   make(map[int64]dtos.Order, 0),
	}
}

func (c *CatalogServiceImpl) CreateProduct(product dtos.Product) (int64, error) {
	if _, ok := c.Products[product.ProdId]; ok {
		err := errors.New("product already exist")
		return 0, err
	}

	productId += 1
	product.ProdId = productId
	c.Products[productId] = product
	fmt.Println(productId)
	fmt.Println(c.Products[productId])
	return productId, nil
}

func (c *CatalogServiceImpl) UpdateProduct(product dtos.Product) error {
	if _, ok := c.Products[product.ProdId]; !ok {
		err := errors.New("product does not exist")
		return err
	}
	c.Products[product.ProdId] = product
	return nil
}

func (c *CatalogServiceImpl) AddProductQty(product *dtos.Product, qty int64) error {
	if _, ok := c.Products[product.ProdId]; !ok {
		err := errors.New("product does not exist")
		return err
	}
	prod := c.Products[product.ProdId]
	prod.QuantityAvailable += qty
	c.Products[product.ProdId] = prod
	return nil
}

func (c *CatalogServiceImpl) DecreaseProductQty(product dtos.Product, qty int64) error {
	if _, ok := c.Products[product.ProdId]; !ok {
		err := errors.New("product does not exist")
		return err
	}
	prod := c.Products[product.ProdId]
	if prod.QuantityAvailable < qty {
		err := errors.New("quantity given to reduce is more that available quantity")
		return err
	}
	prod.QuantityAvailable -= qty
	c.Products[product.ProdId] = prod
	return nil
}

func (c *CatalogServiceImpl) BuyProduct(productId int64, qty int64, email string) (int64, error) {
	if _, ok := c.Products[productId]; !ok {
		err := errors.New("product does not exist")
		return 0, err
	}
	prod := c.Products[productId]
	if prod.QuantityAvailable < qty {
		err := errors.New("quantity left in stock is: " + strconv.Itoa(int(prod.QuantityAvailable)))
		return 0, err
	}

	err := c.DecreaseProductQty(prod, qty)
	if err != nil {
		return 0, err
	}
	orderId += 1
	order := dtos.Order{
		OrderId:    orderId,
		ProdId:     prod.ProdId,
		Quantity:   qty,
		UserEmail:  email,
		TotalPrice: float64(qty) * prod.Price,
		Status:     constants.StatusA,
	}

	c.Orders[orderId] = order
	return orderId, nil
}

func (c *CatalogServiceImpl) UpdateStatus(orderId int64, status string, email string) error {
	if _, ok := c.Orders[orderId]; !ok {
		err := errors.New("order does not exist")
		return err
	}
	order := c.Orders[orderId]
	prod := c.Products[order.ProdId]
	if prod.Seller.Email != email {
		err := errors.New("seller doesnt match")
		return err
	}

	order.Status = status
	c.Orders[orderId] = order
	return nil
}

func (c *CatalogServiceImpl) GetProductById(ProdId int64) (dtos.Product, error) {
	var prod dtos.Product
	if _, ok := c.Products[ProdId]; !ok {
		err := errors.New("product does not exist")
		return prod, err
	}
	prod = c.Products[ProdId]

	return prod, nil
}

func (c *CatalogServiceImpl) GetProducts(filter string) (model.GetProductsResponse, error) {
	var products []model.ProductUpdate
	for _, item := range c.Products {
		var product model.ProductUpdate
		product = model.ProductUpdate{
			ProdId:   item.ProdId,
			Quantity: item.QuantityAvailable,
			Type:     item.Type,
			Price:    item.Price,
			Seller:   model.Seller{Email: item.Seller.Email},
		}
		products = append(products, product)

	}

	return model.GetProductsResponse{Products: products}, nil
}

func (c *CatalogServiceImpl) GetOrders() (model.GetOrderResponse, error) {
	var orders []model.Order
	for _, item := range c.Orders {
		var order model.Order
		order = model.Order{
			ProdId:     item.ProdId,
			Quantity:   item.Quantity,
			OrderId:    item.OrderId,
			TotalPrice: item.TotalPrice,
			Status:     item.Status,
		}
		orders = append(orders, order)

	}

	return model.GetOrderResponse{Orders: orders}, nil
}

func (c *CatalogServiceImpl) GetOrderByID(orderId int64) (model.Order, error) {
	var order model.Order
	for _, item := range c.Orders {
		if item.OrderId == orderId {
			order = model.Order{
				ProdId:     item.ProdId,
				Quantity:   item.Quantity,
				OrderId:    item.OrderId,
				TotalPrice: item.TotalPrice,
				Status:     item.Status,
			}
			break
		}
	}

	return order, nil
}
