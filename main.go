package main

import (
	"connect_db_exc/config"
	"connect_db_exc/entities"
	"connect_db_exc/models"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func main() {
	DemoFindAllViewsAgr()
}

func DemoFindOne() {
	db, err := config.GetMongoDB()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			DB:         db,
			Collection: "product",
		}
		product, err := productModel.FindProduct("61128833ac7c65496003b98f")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(product)
			fmt.Println("Product Info")
			fmt.Println("ID:", product.ID.Hex())
			fmt.Println("Name:", product.Name)
			fmt.Println("Price:", product.Price)
			fmt.Println("Quantity:", product.Quantity)
			fmt.Println("Status:", product.Status)
		}
	}
}

func DemoFindAll() {
	db, err := config.GetMongoDB()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			DB:         db,
			Collection: "product",
		}
		products, err := productModel.FindAllProducts()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(products)
			fmt.Println("Product List")
			for _, product := range products {
				fmt.Println("ID:", product.ID.Hex())
				fmt.Println("Name:", product.Name)
				fmt.Println("Price:", product.Price)
				fmt.Println("Quantity:", product.Quantity)
				fmt.Println("Status:", product.Status)
				fmt.Println("--------------------------------------")
			}
		}
	}
}

func DemoFindAllViewsAgr() {
	db, err := config.GetMongoDB()
	if err != nil {
		fmt.Println(err)
	} else {
		viewModel := models.ViewModel{
			DB:         db,
			Collection: "View",
		}
		views, err := viewModel.FindAllViews()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(views)
			fmt.Println("View List")
			for _, view := range views {
				fmt.Println(view)
				fmt.Println("--------------------------------------")
			}
		}
	}
}

func DemoCreateView() {
	db, err := config.GetMongoDB()
	if err != nil {
		fmt.Println(err)
	} else {
		viewModel := models.ViewModel{
			DB:         db,
			Collection: "View",
		}
		view := entities.View{
			ID:     bson.NewObjectId(),
			Name:   "John",
			UserID: bson.ObjectIdHex("6111462de375d4df10a8e59e"),
			Date:   "11-08-2021",
			Count:  11,
		}
		err := viewModel.CreateView(&view)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(view)
		}
	}
}

func DemoCreate() {
	db, err := config.GetMongoDB()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			DB:         db,
			Collection: "product",
		}
		product := entities.Product{
			ID:       bson.NewObjectId(),
			Name:     "abc",
			Price:    7.77,
			Quantity: 777,
			Status:   true,
		}
		err := productModel.CreateProduct(&product)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(product)
		}
	}
}

func DemoUpdate() {
	db, err := config.GetMongoDB()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			DB:         db,
			Collection: "product",
		}
		product, err := productModel.FindProduct("6112dc3057f2c64934f50d7f")
		if err != nil {
			fmt.Println(err)
		} else {
			product.Name = "Valera"
			product.Price = 777.777
			product.Quantity = 1111
			product.Status = true
			err := productModel.UpdateProduct(&product)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(product)
			}
		}
	}
}

func DemoDelete() {
	db, err := config.GetMongoDB()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			DB:         db,
			Collection: "product",
		}
		product, err := productModel.FindProduct("6112dfc023866b699083481d")
		if err != nil {
			fmt.Println(err)
		} else {
			err := productModel.DeleteProduct(product)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Done! User: %s has been deleted!\n", product.ID)
			}
		}
	}
}

func DemoUpdateView() {
	db, err := config.GetMongoDB()
	if err != nil {
		fmt.Println(err)
	} else {
		viewModel := models.ViewModel{
			DB:         db,
			Collection: "View",
		}
		clickCount := 10
		view, err := viewModel.FindView("611145ce0edd2cbb8f42c85e")
		date := time.Now()
		if err != nil {
			fmt.Println(err)
		} else {
			view.Name = "golang"
			view.Date = date.Format("01-02-2006")
			view.Count = view.Count + clickCount
			err := viewModel.UpdateView(view)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(view)
			}
		}
	}
}
