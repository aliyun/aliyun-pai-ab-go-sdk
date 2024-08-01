package api

import (
	paiabtestClient "github.com/alibabacloud-go/paiabtest-20240119/client"
)

// APIClient manages communication with the PAI-AB Server API API v0.0.1
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	*paiabtestClient.Client
	common        service // Reuse a single struct instead of allocating one for each service on the heap.
	configuration *Configuration
	// API Services

	CrowdApi *CrowdApiService

	DomainApi *DomainApiService

	ExperimentApi *ExperimentApiService

	ExperimentVersionApi *ExperimentVersionApiService

	FeatureApi *FeatureApiService

	LayerApi *LayerApiService

	ProjectApi *ProjectApiService
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(config *Configuration) (*APIClient, error) {
	client, err := paiabtestClient.NewClient(config.GetConfig())
	if err != nil {
		return nil, err
	}
	c := &APIClient{
		Client:        client,
		configuration: config,
	}
	c.common.client = c

	// API Services
	c.CrowdApi = (*CrowdApiService)(&c.common)
	c.DomainApi = (*DomainApiService)(&c.common)
	c.ExperimentApi = (*ExperimentApiService)(&c.common)
	c.ExperimentVersionApi = (*ExperimentVersionApiService)(&c.common)
	c.FeatureApi = (*FeatureApiService)(&c.common)
	c.LayerApi = (*LayerApiService)(&c.common)
	c.ProjectApi = (*ProjectApiService)(&c.common)

	return c, nil
}
