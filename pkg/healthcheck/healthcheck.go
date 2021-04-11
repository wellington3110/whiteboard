package healthcheck

import "net/http"

const Resource = "/ping"

type RouteMapper func(pattern string, handlerFunc http.HandlerFunc)

func Enable(router RouteMapper) {
	router(Resource, func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(200)
	})
}
