package errorrender

import (
	"github.com/go-chi/render"
	"net/http"
)

type HTTPStatusCode interface {
	HTTPStatusCode() int
}

func GetStatusCodeOr(err error, fallback int) int {
	apiError, ok := err.(HTTPStatusCode)
	if !ok {
		return fallback
	}
	return apiError.HTTPStatusCode()
}

func Render(w http.ResponseWriter, r *http.Request, err error) {
	render.Status(r, GetStatusCodeOr(err, http.StatusInternalServerError))
	render.JSON(w, r, err)
}
