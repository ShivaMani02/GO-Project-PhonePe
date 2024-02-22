package model

type ProductCreate struct {
	Quantity int64   `json:"quantity" validate:"required"`
	Price    float64 `json:"price" validate:"required"`
	Type     string  `json:"type" validate:"required"`
	Seller   Seller  `json:"seller"`
}

type ProductUpdate struct {
	ProdId   int64   `json:"prod_id" validate:"required"`
	Quantity int64   `json:"quantity"`
	Price    float64 `json:"price" `
	Type     string  `json:"type"`
	Seller   Seller  `json:"seller" validate:"required"`
}

type Order struct {
	OrderId    int64   `json:"order_id"`
	ProdId     int64   `json:"prod_id"`
	Quantity   int64   `json:"quantity"`
	UserEmail  string  `json:"user_email"`
	TotalPrice float64 `json:"total_price"`
	Status     string  `json:"status"`
}

type Seller struct {
	Email string `json:"email" validate:"required"`
}
