package distsdk

import (
	"github.com/pkg/errors"
)

// client represents an API client struct.
// Any information needed to authenticate/configure a distributor should
// be stored here. This includes IDs, Usernames, Regions, etc.
type client struct {
	ID     string // Example
	Secret string // Example
}

// NewClient returns an APIClient.
// It performs tasks such as API authentication, credentials validation, etc.
func NewClient(opts NewClientOpts) (APIClient, error) {
	if err := opts.validate(); err != nil {
		return nil, errors.Wrap(err, "opts.validate")
	}

	c := &client{
		// TODO Add client details
	}

	c.authenticate()

	return c, nil
}

type NewClientOpts struct {
	ID     string // Example
	Secret string // Example
}

func (opts NewClientOpts) validate() error {
	// TODO validate credentials
	return nil
}
