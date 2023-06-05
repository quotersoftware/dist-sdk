package distsdk

// authenticate is a util func to authenticate an API client.
// This authentication should happen when creating an APIClient (on NewClient).
// The arguments/returns of this function can (and probably should) be modified
// as it will depend on each API.
func (c *client) authenticate() error {
	// TODO add API authentication calls
	return nil
}
