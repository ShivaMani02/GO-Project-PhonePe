package dtos

type Seller struct {
	Name          string
	Email         string
	Password      string
	IsLogin       bool
	SellsProducts []int64 //stores product ids that selled by the seller
	Orders        []int64 //stores order ids of selled products
}
