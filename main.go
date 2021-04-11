package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/wellington3110/whiteboard/pkg/http/middleware"
	"io/ioutil"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Instrument)
	r.Handle("/metrics", promhttp.Handler())
	r.Post("/log", logEndpoint())
	r.Get("/user/{userId}", userEndpoint())
	_ = http.ListenAndServe(":8080", r)
}

func logEndpoint() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		bytes, _ := ioutil.ReadAll(r.Body)
		fmt.Println(string(bytes))
	}
}

func userEndpoint() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := chi.URLParam(r, "userId")
		if userID == "500" {
			w.WriteHeader(500)
			return
		}
		render.Status(r, 200)
		render.JSON(w, r, map[string]string{
			"userId": userID,
		})
	}
}