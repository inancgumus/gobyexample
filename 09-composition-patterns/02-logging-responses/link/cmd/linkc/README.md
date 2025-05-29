# Exercise: Link Client

Implement an HTTP client for the URL shortener server.

## Tasks
- Implement a client for the server's all endpoints.
- The client should be able to shorten and resolve URLs by sending HTTP requests to the server.

## Bonus tasks
- Test the client.
- Add a timeout option and support.
- Add a retry mechanism.

## Proposed API

```go
package rest

// Client is a link management services client.
type Client struct {
	client *http.Client
}

// Shorten shortens the given URL and returns a [link.Link].
func (c *Client) Shorten(_ context.Context, url link.Link) (link.Key, error) {
	return link.Key{}, nil
}

// Resolve resolves the given short URL and returns a [link.Link].
func (c *Client) Resolve(_ context.Context, key link.Key) (link.Link, error) {
	return link.Link{}, nil
}

// Health checks the health of the server.
func (c *Client) Health() error {
	return nil
}
```

# Exercise: Link CLI Client

Create a CLI program called `linkc` that uses the [Link Client](#link-client):
- `linkc shorten <url> # shortens the given URL`
- `linkc resolve <key> # resolves the short URL`
- `linkc health   # checks the health of the server`

Optional flags:

- `-addr`: The address of the server. Default: localhost:8080
- `-timeout`: The timeout for the client. Default: 10 seconds
- `-retry`: The number of retries for the client. Default: 3
- `-retry-wait`: The wait time between retries. Default: 1 second
- `-log`: The log level. Default: info
