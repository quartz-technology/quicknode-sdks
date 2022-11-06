package client

// Opt is a setter for any further client option.
type Opt func(*Client)

func WithAPIKey(key string) Opt {
	return func(c *Client) {
		c.apiKey = key
	}
}
