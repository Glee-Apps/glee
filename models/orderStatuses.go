package models

import (
	"fmt"

	u "github.com/Manuhmutua/glee/utils"

	"github.com/jinzhu/gorm"
)

type OrderStatus struct {
	gorm.Model
	Name string `json:"name"`
}

/*
 This struct function validate the required parameters sent through the http request body
returns message and true if the requirement is met
*/
func (orderStatus *OrderStatus) Validate() (map[string]interface{}, bool) {

	if orderStatus.Name == "" {
		return u.Message(false, "OrderStatus name should be on the payload"), false
	}
	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (orderStatus *OrderStatus) Create() map[string]interface{} {

	if resp, ok := orderStatus.Validate(); !ok {
		return resp
	}

	GetDB().Create(orderStatus)

	resp := u.Message(true, "success")
	resp["order_status"] = orderStatus
	return resp
}

func GetOrderStatuses() []*OrderStatus {

	orderStatuses := make([]*OrderStatus, 0)
	err := GetDB().Table("order_statuses").Find(&orderStatuses).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return orderStatuses
}
