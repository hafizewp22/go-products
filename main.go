package main

import (
	"log"
	"net/http"

	"github.com/hafizewp22/go-products/config"
	"github.com/hafizewp22/go-products/controllers/categorycontroller"
	"github.com/hafizewp22/go-products/controllers/homecontroller"
	"github.com/hafizewp22/go-products/controllers/productcontroller"
)

func main() {
	// 	Database connection
	config.ConnectDB()

	// Routes
	// 1.Homepage
	http.HandleFunc("/", homecontroller.Welcome)

	// 2. Category
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	// 3. Products
	http.HandleFunc("/products", productcontroller.Index)
	http.HandleFunc("/products/add", productcontroller.Add)
	http.HandleFunc("/products/detail", productcontroller.Detail)
	http.HandleFunc("/products/edit", productcontroller.Edit)
	http.HandleFunc("/products/delete", productcontroller.Delete)

	// Run server
	log.Println("Server running on port: 9090")
	log.Fatal(http.ListenAndServe(":9090", nil))
}
