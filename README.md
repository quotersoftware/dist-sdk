# Distributor Integrations SDK Template

```go
type APIClient interface {
    Search(ctx context.Context, opts SearchOpts) (*SearchResponse, error)
}

type client struct {
    ID     string  // Example
    Secret string  // Example
}

type NewClientOpts struct {
    ID     string  // Example
    Secret string  // Example
}

func NewClient(opts NewClientOpts) (APIClient, error) {
    if err := opts.Validate(); err != nil {
	return errors.Wrap(err, "opts.Validate")
    }

    // ...

    return &client{}, error
}


func (c *client) Search(ctx context.Context, opts SearchOpts) (*SearchResponse, error) {
    if err := opts.Validate(); err != nil {
	return errors.Wrap(err, "opts.Validate")
    }


    // ...

    return &SearchResponse{}, error
}
```


