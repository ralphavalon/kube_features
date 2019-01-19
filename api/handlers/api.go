package api

import (
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

var err error

// HeartbeatResponse :: Struct for Healthcheck
type HeartbeatResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

// HealthCheck :: Simple healthcheck endpoint that returns HTTP 200
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	jsoniter.NewEncoder(w).Encode(HeartbeatResponse{Status: "OK", Code: 200})
}
