package api

// HTTP status code and status
// swagger:response healthCheckResponse
type swaggerHealthCheckResponse struct {
	// in:body
	Body HealthCheckResponse
}

// New product request
// swagger:parameters productRequest
type swaggerProductRequest struct {
	// in:body
	Body ProductRequest
}

// Product Response
// swagger:response productResponse
type swaggerProductResponse struct {
	// in:body
	Body ProductResponse
}
