package link

// ---------------------------------------------------------
// Exercise: Add an HTTP client for the URL shortener server.
//
// Explanation:
// You need to implement a client for the URL shortener server. The server is
// already implemented in the link package. The server has two endpoints:
// one for shortening URLs and the other for resolving the short URLs. The
// server also has a health check endpoint. You need to implement a client for
// these endpoints. The client should be able to shorten and resolve URLs using
// the server. The client should also be able to check the health of the server.
//
// You can see the client's proposed API in this file. You need to implement
// the client using the standard library's net/http package.
// ---------------------------------------------------------

// ---------------------------------------------------------
// Exercise: Use the short client you implemented to shorten and resolve URLs.
//
// Explanation:
// Create a program called short that uses the short client you implemented
// to shorten and resolve URLs. Implement the following commands for the shortc:
//
// 	- short -s <url> # shortens the given URL
// 	- short -r <key> # resolves the short URL
// 	- short health   # checks the health of the server
//
// Optional flags:
// - -addr: The address of the server. Default: localhost:8080
// - -timeout: The timeout for the client. Default: 10 seconds
// - -retry: The number of retries for the client. Default: 3
// - -retry-wait: The wait time between retries. Default: 1 second
// - -log: The log level. Default: info
// ---------------------------------------------------------

// ---------------------------------------------------------
// Bonus: Test the client. Reach at least 80% test coverage.
// Bonus: Add a timeout to the client.
// Bonus: Add a retry mechanism to the client.
// ---------------------------------------------------------

// ---------------------------------------------------------
// Warning: Avoid using the server's HTTP handler functions in the client.
// These exercises are about implementing a client for the server using the
// standard library's net/http package.
// ---------------------------------------------------------

// Client is a URL short HTTP client.
type Client struct {
	// *http.Client
}

// NewClient creates and returns a new Client.
func NewClient() *Client {
	return nil
}

// Shorten shortens the given URL and returns a [Link].
func (c *Client) Shorten(url string) (Link, error) {
	return Link{}, nil
}

// Resolve resolves the given short URL and returns a [Link].
func (c *Client) Resolve(key string) (Link, error) {
	return Link{}, nil
}

// Health checks the health of the server.
func (c *Client) Health() error {
	return nil
}
