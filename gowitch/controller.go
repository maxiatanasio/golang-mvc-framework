package gowitch

import "net/http"

type ControllerInterface interface {
	ValidateInputs(r *http.Request) bool
	Handle(w http.ResponseWriter, r *http.Request) interface{}
}

type Controller struct {
	Name string
}
