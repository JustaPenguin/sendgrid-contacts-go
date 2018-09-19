package contacts

import (
	"go.uber.org/ratelimit"
	"net/http"
	"os"
)

var client *Client

func init() {
	client = New(os.Getenv("SENDGRID_APIKEY"))
	client.HTTPClient = &http.Client{
		Transport: &RateLimitingTransport{},
	}
}

type RateLimitingTransport struct{}

var rateLimiter = ratelimit.New(1)

func (x *RateLimitingTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	rateLimiter.Take()

	return http.DefaultTransport.RoundTrip(r)
}
