package client

import (
	"context"
	"net/http"

	"github.com/Khan/genqlient/graphql"
)

// Client is a Graphql Client to Icy Tools GraphQL API.
type Client struct {
	// gql is the base GraphQL Client
	gql graphql.Client

	// apiKey is Icy Tools GraphQL API key.
	// It is used to extend the rate limit.
	apiKey string
}

// New initializes a ready to use Client.
func New(opts ...Opt) *Client {
	c := &Client{}

	for _, opt := range opts {
		opt(c)
	}

	// Initialize Graphql Client
	c.gql = graphql.NewClient(
		Endpoint,
		&http.Client{
			Transport: &authTransport{
				wrapped: http.DefaultTransport,
				key:     c.apiKey,
			},
		})

	return c
}

// MakeRequest executes a raw GraphQL query
func (c *Client) MakeRequest(ctx context.Context, req *graphql.Request, res *graphql.Response) error {
	return c.gql.MakeRequest(ctx, req, res)
}
