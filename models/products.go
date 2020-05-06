package models

import (
	"fmt"

	u "github.com/Manuhmutua/glee/utils"

	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Name              string  `json:"name"`
	Description       string  `json:"description"`
	Quantity          int     `json:"quantity"`
	Cost              int     `json:"cost"`
	Images            []Image `json:"images"`
	UserId            uint    `json:"user_id"`
	ProductCategoryId int     `json:"category_id"`
}

type Image struct {
	gorm.Model
	ProductId uint   `json:"product_id"`
	Url       string `json:"url"`
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

	product := &Product{}
	err := GetDB().Table("products").Where("id = ?", id).First(product).Error
	if err != nil {
		return nil
	}

	images := make([]Image, 0)
	err = GetDB().Table("images").Where("product_id = ?", product.ID).Find(&images).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	product.Images = images

	return product
}

func GetProducts(user uint) []*Product {

	products := make([]*Product, 0)
	err := GetDB().Table("products").Where("user_id = ?", user).Find(&products).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	for _, element := range products {
		images := make([]Image, 0)
		err := GetDB().Table("images").Where("product_id = ?", element.ID).Find(&images).Error
		if err != nil {
			fmt.Println(err)
			return nil
		}
		element.Images = images
	}

	return products
}
