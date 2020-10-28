package api

import (
	"fmt"
	"github.com/davidji99/simpleresty"
	"github.com/go-resty/resty/v2"
	"net/http"
	"strconv"
	"sync"
	"time"
)

const (
	// DefaultAPIBaseURL is the base URL.
	DefaultAPIBaseURL = "https://app.asana.com/api/1.0"

	// DefaultUserAgent is the user agent used when making API calls.
	DefaultUserAgent = "asana-go"

	// DefaultContentTypeHeader is the default  Content-Type header.
	DefaultContentTypeHeader = "application/json"

	// DefaultAcceptHeader is the default Accept header.
	DefaultAcceptHeader = "application/json"

	// DefaultMaxRetryCount is the number of times to retry the request.
	DefaultMaxRetryCount = 3

	// DefaultMaxRetryWaitTime is the number of seconds a retry can wait for.
	DefaultMaxRetryWaitTime = 300 * time.Second
)

// A Client manages communication with the Asana API.
type Client struct {
	// clientMu protects the client during calls that modify the CheckRedirect func.
	clientMu sync.Mutex

	// HTTP client used to communicate with the API.
	http *simpleresty.Client

	// baseURL for API. No trailing slashes.
	baseURL string

	// Reuse a single struct instead of allocating one for each service on the heap.
	common service

	// User agent used when communicating with the Rollbar API.
	userAgent string

	// Custom HTTPHeaders
	customHTTPHeaders map[string]string

	// Access token
	accessToken string

	// Services used for talking to different parts of the Asana API.
	Projects *ProjectsService
}

// service represents the API service client.
type service struct {
	client *Client
}

// CommonResourceFields represents common fields returned by all API resources.
type CommonResourceFields struct {
	// Globally unique identifier of the resource, as a string.
	GID *string `json:"gid,omitempty"`

	// The base type of this resource.
	ResourceType *string `json:"resource_type,omitempty"`
}

// New constructs a new client to interact with the API using a project and/or account access token.
func New(opts ...Option) (*Client, error) {
	// Construct new client.
	c := &Client{
		http:              simpleresty.New(),
		baseURL:           DefaultAPIBaseURL,
		userAgent:         DefaultUserAgent,
		customHTTPHeaders: map[string]string{},
		accessToken:       "",
	}

	// Set Retries
	c.http.SetRetryCount(DefaultMaxRetryCount)
	c.http.SetRetryMaxWaitTime(DefaultMaxRetryWaitTime)
	c.http.AddRetryCondition(
		func(r *resty.Response, err error) bool {
			// Set the retry wait time using the value obtained from the ['Retry-After'] header.
			retryAfterValue, _ := strconv.Atoi(r.Header().Get("Retry-After"))
			c.http.SetRetryWaitTime(time.Duration(retryAfterValue) * time.Second)

			return r.StatusCode() == http.StatusTooManyRequests
		},
	)

	// Define any user custom Client settings
	if optErr := c.parseOptions(opts...); optErr != nil {
		return nil, optErr
	}

	// Setup the client with default settings
	c.setupClient()

	// Inject services
	c.injectServices()

	return c, nil
}

// injectServices adds the services to the client.
func (c *Client) injectServices() {
	c.common.client = c
	c.Projects = (*ProjectsService)(&c.common)
}

// setupClient sets common headers and other configurations.
func (c *Client) setupClient() {
	// Set Base URL
	c.http.SetBaseURL(c.baseURL)

	/*
		We aren't setting an authentication header initially here as certain API resources require specific access_tokens.
		Per Rollbar API documentation, each individual resource will set the access_token parameter when constructing
		the full API endpoint URL.
	*/
	c.http.SetHeader("Content-type", DefaultContentTypeHeader).
		SetHeader("Accept", DefaultAcceptHeader).
		SetHeader("User-Agent", c.userAgent).
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", c.accessToken)).
		SetTimeout(1 * time.Minute).
		SetAllowGetMethodPayload(true)

	// Set additional headers
	if c.customHTTPHeaders != nil {
		c.http.SetHeaders(c.customHTTPHeaders)
	}
}

// parseOptions parses the supplied options functions and returns a configured *Client instance.
func (c *Client) parseOptions(opts ...Option) error {
	// Range over each options function and apply it to our API type to
	// configure it. Options functions are applied in order, with any
	// conflicting options overriding earlier calls.
	for _, option := range opts {
		err := option(c)
		if err != nil {
			return err
		}
	}

	return nil
}
