package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Manuhmutua/glee/app"
	"github.com/Manuhmutua/glee/controllers"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/products/new", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/api/products/category/new", controllers.CreateProductCategory).Methods("POST")
	router.HandleFunc("/api/products/category", controllers.GetProductCategories).Methods("GET")
	router.HandleFunc("/api/me/products", controllers.GetProductFor).Methods("GET")
	router.HandleFunc("/api/orders/new", controllers.CreateOrder).Methods("POST")
	router.HandleFunc("/api/orders/update/{id}", controllers.UpdateOrder).Methods("POST")
	router.HandleFunc("/api/me/orders", controllers.GetOrdersFor).Methods("GET")
	router.HandleFunc("/api/orders/new/status", controllers.CreateOrderStatus).Methods("POST")
	router.HandleFunc("/api/orders", controllers.GetOrderStatuses).Methods("GET")
	router.HandleFunc("/api/user/{id}", controllers.GetUserById).Methods("GET")

	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	//router.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
