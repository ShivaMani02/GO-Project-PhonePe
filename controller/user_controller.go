package controller

import (
	"ProjectPhonePay/model"
	"ProjectPhonePay/service"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type UserController struct {
	UserService service.IUserService
}

func NewUserController(user service.IUserService) UserController {
	return UserController{
		UserService: user,
	}
}

func (c *UserController) UserRegister(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newUser model.UserRegisterReq

	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	validate := validator.New()
	err = validate.Struct(newUser)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	err = c.UserService.UserRegistration(newUser)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err.Error())

		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode("User Register Successfully")
	return
}

func (c *UserController) UserLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var login model.UserLoginReq
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	validate := validator.New()
	err = validate.Struct(login)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	err = c.UserService.UserLogin(login)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err.Error())

		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode("User Successfully LoggedIn")

	return
}

func (c *UserController) UserLogout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var logout model.UserLogoutReq
	err := json.NewDecoder(r.Body).Decode(&logout)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	validate := validator.New()
	err = validate.Struct(logout)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	err = c.UserService.UserLogout(logout)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err.Error())

		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode("User Successfully LoggedOut")

	return
}

func (c *UserController) GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var filter model.GetProducts
	err := json.NewDecoder(r.Body).Decode(&filter)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	validate := validator.New()
	err = validate.Struct(filter)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	products, err := c.UserService.GetProducts(filter.Email, filter.Filter)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err.Error())

		return
	}
	w.WriteHeader(200)
	b, _ := json.Marshal(products)
	w.Write(b)

	return
}

func (c *UserController) GetOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var filter model.GetProducts
	err := json.NewDecoder(r.Body).Decode(&filter)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	validate := validator.New()
	err = validate.Struct(filter)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	products, err := c.UserService.GetOrders(filter.Email)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err.Error())

		return
	}
	w.WriteHeader(200)
	b, _ := json.Marshal(products)
	w.Write(b)

	return
}

func (c *UserController) Checkout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var filter model.GetProducts
	err := json.NewDecoder(r.Body).Decode(&filter)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	validate := validator.New()
	err = validate.Struct(filter)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	strBuy, strNotToAbleBuy, err := c.UserService.Checkout(filter.Email)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err.Error())

		return
	}
	if strNotToAbleBuy != "" {
		strBuy = strBuy + " Not Able to Buy Orders with order (ids: error)" + strNotToAbleBuy
	}

	w.WriteHeader(200)
	b, _ := json.Marshal(strBuy)
	w.Write(b)

	return
}

func (c *UserController) AddToCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req model.AddToCartReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	validate := validator.New()
	err = validate.Struct(req)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	err = c.UserService.AddToCart(req.Email, req.ProdId, req.Quantity)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err.Error())

		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode("Product Added to cart")

	return
}

func (c *UserController) GetAllUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users := c.UserService.GetAllUser()
	b, _ := json.Marshal(users)
	w.WriteHeader(200)
	w.Write(b)
	json.NewEncoder(w).Encode("All User Printed")

	return
}
