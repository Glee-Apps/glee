package models

import (
	"fmt"

	u "github.com/Manuhmutua/glee/utils"

	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
	Cost        int     `json:"cost"`
	Images      []Image `json:"images"`
	UserId      uint    `json:"user_id"` //The user that this contact belongs to
}

type Image struct {
	gorm.Model
	Url string `json:"url"`
}

/*
 This struct function validate the required parameters sent through the http request body
returns message and true if the requirement is met
*/
func (product *Product) Validate() (map[string]interface{}, bool) {

	if product.Name == "" {
		return u.Message(false, "Product name should be on the payload"), false
	}

	if product.Description == "" {
		return u.Message(false, "Description should be on the payload"), false
	}

	if product.Quantity == 0 {
		return u.Message(false, "Quantity should be on the payload"), false
	}

	if product.Cost == 0 {
		return u.Message(false, "Cost should be on the payload"), false
	}

	if product.Images == nil {
		return u.Message(false, "Images should be on the payload"), false
	}

	if product.UserId <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (product *Product) Create() map[string]interface{} {

	if resp, ok := product.Validate(); !ok {
		return resp
	}

	GetDB().Create(product)

	resp := u.Message(true, "success")
	resp["product"] = product
	return resp
}

func GetProduct(id uint) *Product {

	contact := &Product{}
	err := GetDB().Table("contacts").Where("id = ?", id).First(contact).Error
	if err != nil {
		return nil
	}
	return contact
}

func GetProducts(user uint) []*Product {

	contacts := make([]*Product, 0)
	err := GetDB().Table("contacts").Where("user_id = ?", user).Find(&contacts).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return contacts
}
