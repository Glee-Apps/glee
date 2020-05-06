package controllers

import (
	"encoding/json"
	"github.com/Manuhmutua/glee/models"
	u "github.com/Manuhmutua/glee/utils"
	"github.com/gorilla/mux"
	"net/http"
)

var CreateOrder = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(uint) //Grab the id of the user that send the request
	order := &models.Order{}

	err := json.NewDecoder(r.Body).Decode(order)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	order.UserId = user
	resp := order.Create()
	u.Respond(w, resp)
}

var UpdateOrder = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(uint) //Grab the id of the user that send the request
	order := &models.Order{}

	err := json.NewDecoder(r.Body).Decode(order)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	vars := mux.Vars(r)
	orderId := vars["id"]

	order.UserId = user
	resp := order.Update(orderId)
	u.Respond(w, resp)
}

var GetOrdersFor = func(w http.ResponseWriter, r *http.Request) {

	id := r.Context().Value("user").(uint)
	data := models.GetOrders(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var CreateOrderStatus = func(w http.ResponseWriter, r *http.Request) {

	orderStatus := &models.OrderStatus{}

	err := json.NewDecoder(r.Body).Decode(orderStatus)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := orderStatus.Create()
	u.Respond(w, resp)
}

var GetOrderStatuses = func(w http.ResponseWriter, r *http.Request) {

	data := models.GetOrderStatuses()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
