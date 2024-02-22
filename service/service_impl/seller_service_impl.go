package service_impl

import (
	"ProjectPhonePay/dtos"
	"ProjectPhonePay/model"
	"errors"
)

type SellerServiceImpl struct {
	Sellers        map[string]*dtos.Seller
	CatalogService *CatalogServiceImpl
}

func NewSellerService(catalogService *CatalogServiceImpl) *SellerServiceImpl {
	return &SellerServiceImpl{
		Sellers:        make(map[string]*dtos.Seller, 0),
		CatalogService: catalogService,
	}
}

func (u *SellerServiceImpl) SellerRegistration(seller model.SellerRegisterReq) error {
	if _, ok := u.Sellers[seller.Email]; ok {
		err := errors.New("this email id already in use")
		return err
	}
	u.Sellers[seller.Email] = &dtos.Seller{
		Name:     seller.Name,
		Email:    seller.Email,
		Password: seller.Password,
		IsLogin:  false,
	}

	return nil
}

func (u *SellerServiceImpl) SellerLogin(login model.SellerLoginReq) error {
	if _, ok := u.Sellers[login.Email]; !ok {
		err := errors.New("this email id haven't registered yet")
		return err
	}
	seller := u.Sellers[login.Email]
	if seller.Password != login.Password {
		err := errors.New("incorrect password")
		return err
	}
	if seller.IsLogin == true {
		err := errors.New("seller already loggedIn")
		return err
	}
	seller.IsLogin = true
	u.Sellers[login.Email] = seller
	return nil
}

func (u *SellerServiceImpl) SellerLogout(logout model.SellerLogoutReq) error {
	if _, ok := u.Sellers[logout.Email]; !ok {
		err := errors.New("this email id haven't registered yet")
		return err
	}
	seller := u.Sellers[logout.Email]
	if seller.IsLogin == false {
		err := errors.New("seller not loggedIn yet")
		return err
	}
	seller.IsLogin = false
	u.Sellers[logout.Email] = seller
	return nil
}

func (u *SellerServiceImpl) CreateProduct(email string, product model.ProductCreate) error {
	var err error
	if _, ok := u.Sellers[email]; !ok {
		err = errors.New("seller doesnt exist")
		return err
	}
	if !u.Sellers[email].IsLogin {
		err = errors.New("seller doesnt logged in, please login")
		return err
	}

	prod := &dtos.Product{
		QuantityAvailable: product.Quantity,
		Type:              product.Type,
		Price:             product.Price,
		Seller: dtos.Seller{
			Email: email,
		},
	}
	productId, err = u.CatalogService.CreateProduct(*prod)
	u.Sellers[email].SellsProducts = append(u.Sellers[email].SellsProducts, productId)

	return nil

}

func (u *SellerServiceImpl) UpdateStatus(email string, newStatus string, orderId int64) error {
	var err error
	if _, ok := u.Sellers[email]; !ok {
		err = errors.New("seller doesnt exist")
		return err
	}
	
	err = u.CatalogService.UpdateStatus(orderId, newStatus, email)

	return err

}

func (u *SellerServiceImpl) UpdateProduct(email string, product model.ProductUpdate) error {
	var err error
	if _, ok := u.Sellers[email]; !ok {
		err = errors.New("seller doesnt exist")
		return err
	}
	prod := dtos.Product{
		ProdId:            product.ProdId,
		Price:             product.Price,
		QuantityAvailable: product.Quantity,
		Type:              product.Type,
		Seller: dtos.Seller{
			Email: email,
		},
	}
	err = u.CatalogService.UpdateProduct(prod)
	u.Sellers[email].SellsProducts = append(u.Sellers[email].SellsProducts, productId)

	return err

}

func (u *SellerServiceImpl) AddProductQty(email string, product model.ProductUpdate) error {
	var err error
	if _, ok := u.Sellers[email]; !ok {
		err = errors.New("seller doesnt exist")
		return err
	}

	prod := &dtos.Product{
		ProdId: product.ProdId,
		Seller: dtos.Seller{
			Email: email,
		},
	}

	err = u.CatalogService.AddProductQty(prod, product.Quantity)
	u.Sellers[email].SellsProducts = append(u.Sellers[email].SellsProducts, productId)

	return err

}

func (u *SellerServiceImpl) GetAllSeller() []dtos.Seller {
	var sellers []dtos.Seller
	for email := range u.Sellers {
		sellers = append(sellers, *u.Sellers[email])
	}
	return sellers
}
