package api

import (
	"encoding/json"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

var err error

type ProductRequest struct {
	Name  string
	Price uint
}

type ProductResponse struct {
	ID    string
	Name  string
	Price uint
}

// HeartbeatResponse :: Struct for Healthcheck
type HeartbeatResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

// HealthCheck :: Simple healthcheck endpoint that returns HTTP 200
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	jsoniter.NewEncoder(w).Encode(HeartbeatResponse{Status: "OK", Code: 200})
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
