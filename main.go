package main

import (
	"github.com/wellington3110/whiteboard/internal/domain"
	"net/http"
)

func main() {
	_ = http.ListenAndServe(":8080", domain.Bootstrap())
}