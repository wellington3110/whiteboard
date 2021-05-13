package domain

import (
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	biggestIntEndpoint "github.com/wellington3110/whiteboard/internal/domain/biggestint/endpoint"
	biggestIntService "github.com/wellington3110/whiteboard/internal/domain/biggestint/service"
	"github.com/wellington3110/whiteboard/pkg/healthcheck"
	"github.com/wellington3110/whiteboard/pkg/http/httpmiddleware"
	"github.com/wellington3110/whiteboard/pkg/logger/loggermiddleware"
	"net/http"
)

func Bootstrap() http.Handler {
	router := chi.NewRouter()
	setupMiddleware(router)
	healthcheck.Enable(router.HandleFunc)
	biggestIntEndpoint.BindHTTP(router, biggestIntService.New())
	router.Handle("/metrics", promhttp.Handler())
	return router
}

func setupMiddleware(router *chi.Mux) {
	router.Use(httpmiddleware.RequestID)
	router.Use(httpmiddleware.InstrumentIgnoring(healthcheck.Resource))
	router.Use(loggermiddleware.ContextApplying(loggermiddleware.RequestIDApply))
}