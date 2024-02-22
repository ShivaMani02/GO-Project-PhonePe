package service_impl

import (
	"ProjectPhonePay/dtos"
	"ProjectPhonePay/model"
	"errors"
	"fmt"
	"strconv"
)

type UserServiceImpl struct {
	Users          map[string]*dtos.User
	CatalogService *CatalogServiceImpl
}

func NewUserService(catalogService *CatalogServiceImpl) *UserServiceImpl {
	return &UserServiceImpl{
		Users:          make(map[string]*dtos.User, 0),
		CatalogService: catalogService,
	}
}

func (u *UserServiceImpl) UserRegistration(user model.UserRegisterReq) error {
	if _, ok := u.Users[user.Email]; ok {
		err := errors.New("this email id already in use")
		return err
	}
	u.Users[user.Email] = &dtos.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		IsLogin:  false,
	}

	return nil
}

func (u *UserServiceImpl) UserLogin(login model.UserLoginReq) error {
	if _, ok := u.Users[login.Email]; !ok {
		err := errors.New("this email id haven't registered yet")
		return err
	}
	user := u.Users[login.Email]
	if user.Password != login.Password {
		err := errors.New("incorrect password")
		return err
	}
	if user.IsLogin == true {
		err := errors.New("user already loggedIn")
		return err
	}
	user.IsLogin = true
	u.Users[login.Email] = user
	return nil
}

func (u *UserServiceImpl) UserLogout(logout model.UserLogoutReq) error {
	if _, ok := u.Users[logout.Email]; !ok {
		err := errors.New("this email id haven't registered yet")
		return err
	}
	user := u.Users[logout.Email]
	if user.IsLogin == false {
		err := errors.New("user not loggedIn yet")
		return err
	}
	user.IsLogin = false
	u.Users[logout.Email] = user
	return nil
}

func (u *UserServiceImpl) AddToCart(email string, productId int64, qty int64) error {

	var err error
	if _, ok := u.Users[email]; !ok {
		err = errors.New("user doesnt exist")
		return err
	}

	user := u.Users[email]
	if user.IsLogin == false {
		err = errors.New("user not loggedIn yet")
		return err
	}

	_, err = u.CatalogService.GetProductById(productId)
	if err != nil {
		return err
	}

	cart := dtos.Cart{
		ProdId:   productId,
		Quantity: qty,
	}

	user.Cart = append(user.Cart, cart)
	u.Users[email] = user

	return nil
}

func (u *UserServiceImpl) RemoveToCart(email string, prodId int64, qty int64) error {

	var err error
	if _, ok := u.Users[email]; !ok {
		err = errors.New("user doesnt exist")
		return err
	}

	user := u.Users[email]
	if user.IsLogin == false {
		err = errors.New("user not loggedIn yet")
		return err
	}

	_, err = u.CatalogService.GetProductById(productId)
	if err != nil {
		return err
	}

	for i, cartItem := range u.Users[email].Cart {
		if cartItem.ProdId == prodId {
			if cartItem.Quantity < qty {
				err = errors.New("cart doesn't have enough qty of given item")
				return err
			}
			u.Users[email].Cart[i].Quantity = cartItem.Quantity - qty
			return nil
		}
	}

	err = errors.New("product doesnt exist in cart")
	return err
}

func (u *UserServiceImpl) Checkout(email string) (string, string, error) {
	var err error
	if _, ok := u.Users[email]; !ok {
		err = errors.New("user doesnt exist")
		return "", "", err
	}
	user := u.Users[email]
	if user.IsLogin == false {
		err = errors.New("user not loggedIn yet")
		return "", "", err
	}
	if len(user.Cart) <= 0 {
		err = errors.New("cart empty")
		return "", "", err
	}
	strNotToAbleBuy := ""
	strBuy := "Purchased Orders with order ids: "
	for _, cartItem := range u.Users[email].Cart {
		fmt.Println(cartItem.ProdId)
		orderId, err = u.CatalogService.BuyProduct(cartItem.ProdId, cartItem.Quantity, email)
		if err != nil {
			fmt.Println(err.Error())
			strNotToAbleBuy = strNotToAbleBuy + " (" + strconv.Itoa(int(cartItem.ProdId)) + " : " + err.Error() + ") ,"
		} else {
			u.Users[email].Orders = append(u.Users[email].Orders, orderId)
			strBuy = strBuy + strconv.Itoa(int(cartItem.ProdId)) + " ,"
		}
	}
	if strBuy == "Purchased Orders with order ids: " {
		strBuy = ""
	}

	return strBuy, strNotToAbleBuy, nil
}

func (u *UserServiceImpl) BuyItem(email string, prodId int64, qty int64) error {
	var err error
	if _, ok := u.Users[email]; !ok {
		err = errors.New("user doesnt exist")
		return err
	}

	user := u.Users[email]
	if user.IsLogin == false {
		err = errors.New("user not loggedIn yet")
		return err
	}

	orderId, err = u.CatalogService.BuyProduct(prodId, qty, email)
	if err != nil {
		return err
	}

	u.Users[email].Orders = append(u.Users[email].Orders, orderId)

	return nil
}

func (u *UserServiceImpl) GetProducts(email string, filter string) (model.GetProductsResponse, error) {
	var product model.GetProductsResponse
	var err error
	if _, ok := u.Users[email]; !ok {
		err = errors.New("user doesnt exist")
		return product, err
	}

	user := u.Users[email]
	if user.IsLogin == false {
		err = errors.New("user not loggedIn yet")
		return product, err
	}

	products, err := u.CatalogService.GetProducts(filter)
	if err != nil {
		return product, err
	}

	return products, nil
}

func (u *UserServiceImpl) GetOrders(email string) (model.GetOrderResponse, error) {
	var orders model.GetOrderResponse
	var err error
	if _, ok := u.Users[email]; !ok {
		err = errors.New("user doesnt exist")
		return orders, err
	}

	user := u.Users[email]
	if user.IsLogin == false {
		err = errors.New("user not loggedIn yet")
		return orders, err
	}

	var res []model.Order
	for _, Id := range user.Orders {
		order, err := u.CatalogService.GetOrderByID(Id)
		if err != nil {
			return orders, err
		}
		res = append(res, order)
	}

	orders = model.GetOrderResponse{
		Orders: res,
	}
	if err != nil {
		return orders, err
	}

	return orders, nil
}

func (u *UserServiceImpl) GetAllUser() []dtos.User {
	var users []dtos.User
	for email := range u.Users {
		users = append(users, *u.Users[email])
	}
	return users
}
