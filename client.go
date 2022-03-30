package iracing

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/cookiejar"
	"strconv"
	"time"
)

const (
	Host      = "https://members-ng.iracing.com"
	ImageHost = "https://images-static.iracing.com"
	UserAgent = "iracing-data-go +https://github.com/LeoAdamek/iracing-data-go"
)

var (
	ErrServerError = errors.New("Server Error")
	ErrClientError = errors.New("Client Error")
)

// RateLimiting data
type rateLimiting struct {
	Reset     time.Time
	Remaining uint64
	Total     uint64
}

// iRacing Data API Client
type Client struct {
	inner       *http.Client
	credentials CredentialsProvider
	rateLimiting
	Verbose bool
}

type CacheLink struct {
	URL string `json:"link"`
}

// Create a new client with the given CredentialsProvider
func New(credentials CredentialsProvider) *Client {
	jar, _ := cookiejar.New(nil)

	return &Client{
		inner: &http.Client{
			Jar: jar,
		},
		credentials: credentials,
		Verbose:     false,
	}
}

func (c *Client) do(req *http.Request) (*http.Response, error) {

	req.Header.Set("user-agent", UserAgent)

	res, err := c.inner.Do(req)

	if err != nil {
		return nil, err
	}

	if remaining := res.Header.Get("X-RateLimit-Remaining"); remaining != "" {
		c.rateLimiting.Remaining, err = strconv.ParseUint(remaining, 10, 64)
	}

	if total := res.Header.Get("x-ratelimit-limit"); total != "" {
		c.rateLimiting.Total, err = strconv.ParseUint(total, 10, 64)
	}

	if reset := res.Header.Get("x-ratelimit-reset"); reset != "" {
		ts, _ := strconv.ParseInt(reset, 10, 64)
		c.rateLimiting.Reset = time.Unix(ts, 0)
	}

	if res.StatusCode >= 400 {
		if res.StatusCode >= 500 {
			return nil, ErrServerError
		}
		return nil, ErrClientError
	}

	return res, err
}

func (c *Client) json(method string, url string, body interface{}, into interface{}) error {

	var req *http.Request
	var err error

	if body != nil {
		buffer := new(bytes.Buffer)
		encoder := json.NewEncoder(buffer)

		if err := encoder.Encode(body); err != nil {
			return err
		}
		req, err = http.NewRequest(method, url, buffer)
		req.Header.Set("content-type", "application/json")
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		return err
	}

	res, err := c.do(req)

	if err != nil {
		return err
	}

	if into != nil {
		decoder := json.NewDecoder(res.Body)

		if err := decoder.Decode(&into); err != nil {
			return err
		}
	}

	return nil
}
