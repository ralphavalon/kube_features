package api

import (
	"encoding/json"
	"fmt"
	"kube_features/api/data"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

var err error

// ProductRequest :: Request model for product
type ProductRequest struct {
	// Name of the product
	Name string `json:"name"`
	// Price of the product
	Price uint `json:"price"`
}

// ProductResponse :: Response model for product
type ProductResponse struct {
	// ID of the product
	ID uint `json:"id"`
	// Name of the product
	Name string `json:"name"`
	// Price of the product
	Price uint `json:"price"`
}

// HealthCheckResponse :: Response model for health check
type HealthCheckResponse struct {
	// "OK" or "FAILED"
	Status string `json:"status"`
	// HTTP status code
	Code int `json:"code"`
	// Version
	Version string `json:"version"`
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Println("• Checking for health...")
	jsoniter.NewEncoder(w).Encode(HealthCheckResponse{Status: "OK", Code: 200, Version: "v1"})
}

func CreateProduct(w http.ResponseWriter, request *http.Request) {
	fmt.Println("• Creating product...")
	decoder := json.NewDecoder(request.Body)
	var product ProductRequest
	err := decoder.Decode(&product)
	if err != nil {
		panic(err)
	}
	_, createdProduct, _ := data.CreateProduct(product.Name, product.Price)
	jsoniter.NewEncoder(w).Encode(ProductResponse{ID: createdProduct.Model.ID, Name: createdProduct.Name, Price: createdProduct.Price})
}
