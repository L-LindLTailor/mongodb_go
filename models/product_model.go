package models

import (
	"connect_db_exc/entities"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ProductModel struct {
	DB         *mgo.Database
	Collection string
}

func (productModel ProductModel) FindAllProducts() (products []entities.Product, err error) {
	err = productModel.DB.C(productModel.Collection).Find(bson.M{}).All(&products)
	return
}

func (productModel ProductModel) FindProduct(id string) (product entities.Product, err error) {
	err = productModel.DB.C(productModel.Collection).FindId(bson.ObjectIdHex(id)).One(&product)
	return
}

func (productModel ProductModel) CreateProduct(product *entities.Product) (err error) {
	err = productModel.DB.C(productModel.Collection).Insert(&product)
	return
}

func (productModel ProductModel) UpdateProduct(product *entities.Product) (err error) {
	err = productModel.DB.C(productModel.Collection).UpdateId(product.ID, &product)
	return
}

func (productModel ProductModel) DeleteProduct(product entities.Product) (err error) {
	err = productModel.DB.C(productModel.Collection).Remove(product)
	return
}
