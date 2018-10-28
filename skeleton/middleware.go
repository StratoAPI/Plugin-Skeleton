package skeleton

import "github.com/StratoAPI/Interface/middleware"

type SkeletonMiddleware struct {
}

// Initialize the facade.
func (middleware SkeletonMiddleware) Initialize() error {
	return nil // TODO
}

// Start the facade. Must be a blocking call.
func (middleware SkeletonMiddleware) Start() error {
	return nil // TODO
}

// Graceful stopping of the facade with a 30s timeout.
func (middleware SkeletonMiddleware) Stop() error {
	return nil // TODO
}

// Validate request
// If returns nil, it will pass the request onwards
// If returns a pointer, request will get replied to with provided response
func (middleware SkeletonMiddleware) Request(resource string, headers map[string][]string, config interface{}) *middleware.RequestResponse {
	return nil // TODO
}

// Validate and filter response
// If returns nil, it will pass the request onwards with the returned data
// If returns a pointer, request will get replied to with provided response without data
func (middleware SkeletonMiddleware) Response(resource string, headers map[string][]string, data []map[string]interface{}, config interface{}) ([]map[string]interface{}, *middleware.RequestResponse) {
	return data, nil // TODO
}
