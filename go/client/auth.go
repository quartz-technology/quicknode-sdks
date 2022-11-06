package client

import "net/http"

// authTransport is a wrapper around Client to add authentication header
type authTransport struct {
	wrapped http.RoundTripper
	key     string
}

func (t *authTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("x-api-key", t.key)
	return t.wrapped.RoundTrip(req)
}
