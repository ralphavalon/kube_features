package api

// HTTP status code and status and version
// swagger:response healthCheckResponse
type swaggerHealthCheckResponse struct {
	// in:body
	Body HealthCheckResponse
}

// Current version and called version
// swagger:response versionCheckResponse
type swaggerVersionCheckResponse struct {
	// in:body
	Body VersionCheckResponse
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
