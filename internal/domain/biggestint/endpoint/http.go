package endpoint

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/wellington3110/whiteboard/internal/domain/biggestint"
	"github.com/wellington3110/whiteboard/pkg/http/render/errorrender"
	"github.com/wellington3110/whiteboard/pkg/slice"
	"net/http"
	"strings"
)

type Service interface {
	FindBiggestInt(context.Context, []int) (int, error)
}

type HTTP struct {
	service Service
}

func BindHTTP(router chi.Router, service Service) {
	endpoint := &HTTP{service: service}
	router.Get("/biggest-int", endpoint.FindBiggestInt)
}

func (h *HTTP) FindBiggestInt(w http.ResponseWriter, r *http.Request) {
	ints, err := getInts(r)
	if err != nil {
		errorrender.Render(w, r, err)
		return
	}
	biggestInt, err := h.service.FindBiggestInt(r.Context(), ints)
	if err != nil {
		errorrender.Render(w, r, err)
		return
	}
	render.JSON(w, r, map[string]int{
		"biggestInt": biggestInt,
	})
}

func getInts(r *http.Request) ([]int, error) {
	input := r.URL.Query().Get("ints")
	if input == "" {
		return nil, nil
	}
	supposedInts := slice.TrimValues(strings.Split(input, ","))
	ints, err := slice.FromStringToInt(supposedInts)
	if err != nil {
		return nil, biggestint.NewMustContainOnlyIntsError(input)
	}
	return ints, nil
}
