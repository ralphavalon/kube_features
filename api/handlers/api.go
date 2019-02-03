package api

import (
	"encoding/json"
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
	ID string `json:"id"`
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
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	jsoniter.NewEncoder(w).Encode(HealthCheckResponse{Status: "OK", Code: 200})
}

func CreateProduct(w http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var product ProductRequest
	err := decoder.Decode(&product)
	if err != nil {
		panic(err)
	}
	jsoniter.NewEncoder(w).Encode(ProductResponse{ID: "fakeId", Name: product.Name, Price: product.Price})
}
