package main

import (
	"Golang-Web/config"
	"Golang-Web/controllers/categoriescontroller"
	"Golang-Web/controllers/homecontroller"
	"Golang-Web/controllers/productcontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	// 1. Homepage
	http.HandleFunc("/", homecontroller.Welcome)

	// 2. Category
	http.HandleFunc("/categories", categoriescontroller.Index)
	http.HandleFunc("/categories/add", categoriescontroller.Add)
	http.HandleFunc("/categories/edit", categoriescontroller.Edit)
	http.HandleFunc("/categories/delete", categoriescontroller.Delete)

	// 3. Products
	http.HandleFunc("/product", productcontroller.Index)
	http.HandleFunc("/product/add", productcontroller.Add)
	http.HandleFunc("/product/detail", productcontroller.Detail)
	http.HandleFunc("/product/edit", productcontroller.Edit)
	http.HandleFunc("/product/delete", productcontroller.Delete)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
