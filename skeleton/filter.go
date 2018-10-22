package skeleton

import "github.com/StratoAPI/Interface/filter"

type SkeletonFilter struct {
}

// Initialize the filter.
func (skeleton SkeletonFilter) Initialize() error {
	return nil // TODO
}

// Start the filter. Does not have to be blocking.
func (skeleton SkeletonFilter) Start() error {
	return nil // TODO
}

// Graceful stopping of the filter with a 30s timeout.
func (skeleton SkeletonFilter) Stop() error {
	return nil // TODO
}

// Validate structure for filter validness
func (skeleton SkeletonFilter) ValidateFilter(filter filter.ProcessedFilter) (bool, error) {
	return false, nil // TODO
}

// Create a new instance of the filter
func (skeleton SkeletonFilter) CreateFilter(filter string) (interface{}, error) {
	return nil, nil // TODO
}
