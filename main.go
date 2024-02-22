package main

import (
	"ProjectPhonePay/controller"
	"ProjectPhonePay/service/service_impl"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	catalogService := service_impl.NewCatalogService()
	userService := service_impl.NewUserService(catalogService)
	userController := controller.NewUserController(userService)
	sellerService := service_impl.NewSellerService(catalogService)
	sellerController := controller.NewSellerController(sellerService)
	r := mux.NewRouter()
	r.HandleFunc("/user/register", userController.UserRegister).Methods(http.MethodPost)
	r.HandleFunc("/user/login", userController.UserLogin).Methods(http.MethodPost)
	r.HandleFunc("/user/logout", userController.UserLogout).Methods(http.MethodPost)
	r.HandleFunc("/user/get_all", userController.GetAllUser).Methods(http.MethodGet)
	r.HandleFunc("/user/get_all_product", userController.GetProducts).Methods(http.MethodGet)
	r.HandleFunc("/user/get_all_order", userController.GetOrders).Methods(http.MethodGet)
	r.HandleFunc("/user/add_to_cart", userController.AddToCart).Methods(http.MethodPost)
	r.HandleFunc("/user/checkout", userController.Checkout).Methods(http.MethodPost)

	r.HandleFunc("/seller/register", sellerController.SellerRegister).Methods(http.MethodPost)
	r.HandleFunc("/seller/login", sellerController.SellerLogin).Methods(http.MethodPost)
	r.HandleFunc("/seller/logout", sellerController.SellerLogout).Methods(http.MethodPost)
	r.HandleFunc("/seller/get_all", sellerController.GetAllSeller).Methods(http.MethodGet)
	r.HandleFunc("/seller/new_product", sellerController.CreateProduct).Methods(http.MethodPost)

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
