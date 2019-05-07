package gowitch

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Router struct {
	middlewares []func(next http.Handler) http.Handler
	routes      []Route
}

var muxRouter *mux.Router
var gowitchRouter Router

func RegisterMiddleware(middleware func(next http.Handler) http.Handler) {
	gowitchRouter.middlewares = append(gowitchRouter.middlewares, middleware)
}

func RegisterRoute(route Route) {
	gowitchRouter.routes = append(gowitchRouter.routes, route)
}

func Init() {

	muxRouter = mux.NewRouter().StrictSlash(true)

	for _, middleware := range gowitchRouter.middlewares {
		muxRouter.Use(middleware)
	}

	for _, route := range gowitchRouter.routes {
		addRoute(route)
	}

	log.Fatal(http.ListenAndServe(":8081", muxRouter))

}

func addRoute(route Route) {
	muxRouter.HandleFunc(route.Path, func(w http.ResponseWriter, r *http.Request) {
		var result interface{}
		if !route.Controller.ValidateInputs(r) {
			result = map[string]string{
				"message": "Invalid Inputs",
			}

		} else {
			result = route.Controller.Handle(w, r)
		}
		json.NewEncoder(w).Encode(result)
	}).Methods(route.Method)
}
