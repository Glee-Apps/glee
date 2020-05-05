package models

import (
	"fmt"

	u "github.com/Manuhmutua/glee/utils"

	"github.com/jinzhu/gorm"
)

type ProductCategory struct {
	gorm.Model
	Name string `json:"name"`
}

/*
 This struct function validate the required parameters sent through the http request body
returns message and true if the requirement is met
*/
func (productCategory *ProductCategory) Validate() (map[string]interface{}, bool) {

	if productCategory.Name == "" {
		return u.Message(false, "ProductCategory name should be on the payload"), false
	}
	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (productCategory *ProductCategory) Create() map[string]interface{} {

	if resp, ok := productCategory.Validate(); !ok {
		return resp
	}

	GetDB().Create(productCategory)

	resp := u.Message(true, "success")
	resp["product_category"] = productCategory
	return resp
}

func GetProductCategories() []*ProductCategory {

	productCategories := make([]*ProductCategory, 0)
	err := GetDB().Table("product_categories").Find(&productCategories).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return productCategories
}
