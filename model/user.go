package model

type UserRegisterReq struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Address  string `json:"address"`
}

type UserLoginReq struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserLogoutReq struct {
	Email string `json:"email" validate:"required"`
}

type GetProducts struct {
	Email  string `json:"email" validate:"required"`
	Filter string `json:"filter"`
}

type GetProductsResponse struct {
	Products []ProductUpdate
}

type GetOrderResponse struct {
	Orders []Order
}

type AddToCartReq struct {
	Email    string `json:"email" validate:"required"`
	ProdId   int64  `json:"prod_id" validate:"required"`
	Quantity int64  `json:"quantity" validate:"required"`
}
