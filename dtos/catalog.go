package dtos

type Product struct {
	ProdId            int64
	Price             float64
	QuantityAvailable int64
	Seller            Seller
	Type              string
}

type Order struct {
	OrderId    int64
	ProdId     int64
	Quantity   int64
	UserEmail  string
	TotalPrice float64
	Status     string
}
