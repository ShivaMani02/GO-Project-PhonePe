package dtos

type User struct {
	Name     string
	Email    string
	Password string
	IsLogin  bool
	Orders   []int64 //stores order ids of purchased products
	Cart     []Cart
	Address  string
}

type Cart struct {
	ProdId   int64
	Quantity int64
}
