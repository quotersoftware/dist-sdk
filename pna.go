package distsdk

import (
	"context"

	"github.com/pkg/errors"
)

type APIClient interface {
	Search(ctx context.Context, opts SearchOpts) (*SearchResponse, error)
}

// Search searches products for a certain distributor based on one or multiple identifiers.
func (c *client) Search(ctx context.Context, opts SearchOpts) (*SearchResponse, error) {
	if err := opts.validate(); err != nil {
		return nil, errors.Wrap(err, "opts.validate")
	}

	// TODO implement search calls from the API and build response

	return &SearchResponse{}, nil
}

type SearchOpts struct {
	// UniqueIdentifier represents any sort of unique product identifier for
	// a distributor
	UniqueIdentifier string
}

func (opts SearchOpts) validate() error {
	// TODO validate search params
	return nil
}

type SearchResponse struct {
	Products []Product // Example
}

// Product is an example of a Distributor product
// It should contain product information such as Name, MPN, Availability, etc.
type Product struct {
	Name         string // Example
	MPN          string // Example
	Availability int    // Example
}
