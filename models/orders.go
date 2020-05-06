package models

import (
	"fmt"

	u "github.com/Manuhmutua/glee/utils"

	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	Quantity  int  `json:"quantity"`
	UserId    uint `json:"user_id"`
	ProductId int  `json:"product_id"`
}

/*
 This struct function validate the required parameters sent through the http request body
returns message and true if the requirement is met
*/
func (order *Order) Validate() (map[string]interface{}, bool) {

	if order.Quantity == 0 {
		return u.Message(false, "Quantity should be on the payload"), false
	}

	if order.ProductId <= 0 {
		return u.Message(false, "Product is not recognized"), false
	}

	if order.UserId <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (order *Order) Create() map[string]interface{} {

	if resp, ok := order.Validate(); !ok {
		return resp
	}

	GetDB().Create(order)

	resp := u.Message(true, "success")
	resp["order"] = order
	return resp
}

func (order *Order) Update(id int) map[string]interface{} {

	if resp, ok := order.Validate(); !ok {
		return resp
	}

	GetDB().Update(order).Where("id = ?", id)

	resp := u.Message(true, "success")
	resp["order"] = order
	return resp
}

func GetOrder(id uint) *Order {

	order := &Order{}
	err := GetDB().Table("orders").Where("id = ?", id).First(order).Error
	if err != nil {
		return nil
	}
	return order
}

func GetOrders(user uint) []*Order {

	orders := make([]*Order, 0)
	err := GetDB().Table("orders").Where("user_id = ?", user).Find(&orders).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return orders
}
