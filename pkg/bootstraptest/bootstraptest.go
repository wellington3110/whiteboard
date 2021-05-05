package bootstraptest

import (
	"github.com/h2non/baloo"
	"net/http"
	"net/http/httptest"
)

func New(handler http.Handler) (*httptest.Server, *baloo.Client) {
	server := httptest.NewServer(handler)
	return server, baloo.New(server.URL)
}
