package httpmiddleware

import (
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func RequestID(next http.Handler) http.Handler {
	return middleware.RequestID(next)
}
