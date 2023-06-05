# Distributor Integrations SDK Template
This is a guide for developing a Distributor Integration Go SDK which Quoter can integrate on [Product Cloud](https://quoter.com/features/product-cloud).

## Explanation
This template is an API Wrapper which implements the `Search` function from the `APIClient`. This gives the ability for Quoter APIs to search and display products on [Product Cloud](quoter.com/features/product-cloud).

## Usage
1. Fork this reposistory.
2. Complete and adapt the TODOs on the code for your API needs.
3. Publish a release (on Github) following [semver](https://semver.org/).

## Example
```go
func main() {
	client, err := sdk.NewClient(sdk.NewClientOpts{
		ID:     clientID,
		Secret: clientSecret,
	})
	if err != nil {
	    fmt.Println(err)
	}
	
	resp, err := client.Search(context.Background(), SearchOpts{
            UniqueIdentifier: "product-123",
        })
	if err != nil {
	    fmt.Println(err)
	    return
	}

	fmt.Println(resp)
}
```

## File Structure
```
auth.go   -> API authentication utils 
client.go -> API client initialization and implementation 
pna.go    -> Product and Availability (Search) implementation
```

## Suggestions
Suggestions or faced any issues? Open an issue [here](https://github.com/quotersoftware/dist-sdk/issues).

