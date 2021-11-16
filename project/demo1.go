package project

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Product struct {
	Id          int     `json:"id"`
	ProductName string  `json:"productName"`
	CategoryId  int     `json:"categoryId"`
	UnitPrice   float64 `json:"unitPrice"`
}

type Category struct {
	Id           int    `json:"id"`
	categoryName string `json:"categoryName"`
}

func GetAllProducts() ([]Product, error) {
	response, err := http.Get("http://127.0.0.1:3000/products")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)

	var products []Product
	json.Unmarshal(bodyBytes, &products)
	return products, nil
}

func AddProduct() (Product, error) {
	product := Product{ProductName: "3D Printer", CategoryId: 1, UnitPrice: 200}
	jsonProduct, err := json.Marshal(product)

	response, err := http.Post("http://127.0.0.1:3000/products", "application/json;charset=utf-8", bytes.NewBuffer(jsonProduct))
	if err != nil {
		return Product{}, err
	}
	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)
	var productResponse Product
	json.Unmarshal(bodyBytes, &productResponse)
	fmt.Println("Kaydedildi: ", productResponse)
	return productResponse, nil
}
