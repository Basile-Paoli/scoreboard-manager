package api

import (
	"github.com/Khan/genqlient/graphql"
	"net/http"
)

type transport struct{}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("client-version", "20")
	return http.DefaultTransport.RoundTrip(req)
}

func NewClient() graphql.Client {
	httpClient := &http.Client{
		Transport: &transport{},
	}
	return graphql.NewClient("https://www.start.gg/api/-/gql", httpClient)
}
