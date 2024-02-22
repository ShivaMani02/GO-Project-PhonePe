package model

type SellerRegisterReq struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SellerLoginReq struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SellerLogoutReq struct {
	Email string `json:"email" validate:"required"`
}
