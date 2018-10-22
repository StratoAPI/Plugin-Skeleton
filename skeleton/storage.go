package skeleton

import "github.com/StratoAPI/Interface/filter"

type SkeletonStorage struct {
}

// Initialize the storage.
func (storage SkeletonStorage) Initialize() error {
	return nil // TODO
}

// Start the storage. Must be a blocking call.
func (storage SkeletonStorage) Start() error {
	return nil // TODO
}

// Graceful stopping of the storage with a 30s timeout.
func (storage SkeletonStorage) Stop() error {
	return nil // TODO
}

// Retrieve resources.
func (storage SkeletonStorage) GetResources(resource string, filters []filter.ProcessedFilter) ([]map[string]interface{}, error) {
	return nil, nil // TODO
}

// Create resources.
func (storage SkeletonStorage) CreateResources(resource string, data []map[string]interface{}) error {
	return nil // TODO
}

// Update resources.
func (storage SkeletonStorage) UpdateResources(resource string, data []map[string]interface{}, filters []filter.ProcessedFilter) error {
	return nil // TODO
}

// Delete resources.
func (storage SkeletonStorage) DeleteResources(resource string, filters []filter.ProcessedFilter) error {
	return nil // TODO
}
