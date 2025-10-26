package routes

import "net/http"

type HeaderTransport struct {
	headers map[string]string
	base    http.RoundTripper
}

func (t *HeaderTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	clone := r.Clone(r.Context())

	for k, v := range t.headers {
		if clone.Header.Get(k) == "" {
			clone.Header.Set(k, v)
		}
	}

	return t.base.RoundTrip(clone)
}
