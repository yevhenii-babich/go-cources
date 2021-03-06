/*
 * Scaffold Servise
 *
 * Base for other projects
 *
 * API version: 0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/dev/",
		Index,
	},

	Route{
		"MIdBuzzGet",
		strings.ToUpper("Get"),
		"/dev/{mId}/buzz",
		MIdBuzzGet,
	},

	Route{
		"MIdBuzzIdDelete",
		strings.ToUpper("Delete"),
		"/dev/{mId}/buzz/{Id}",
		MIdBuzzIdDelete,
	},

	Route{
		"MIdBuzzIdGet",
		strings.ToUpper("Get"),
		"/dev/{mId}/buzz/{Id}",
		MIdBuzzIdGet,
	},

	Route{
		"MIdBuzzIdPut",
		strings.ToUpper("Put"),
		"/dev/{mId}/buzz/{Id}",
		MIdBuzzIdPut,
	},

	Route{
		"MIdBuzzPost",
		strings.ToUpper("Post"),
		"/dev/{mId}/buzz",
		MIdBuzzPost,
	},

	Route{
		"CouchGet",
		strings.ToUpper("Get"),
		"/dev/couch",
		CouchGet,
	},

	Route{
		"CouchIdDelete",
		strings.ToUpper("Delete"),
		"/dev/couch/{Id}",
		CouchIdDelete,
	},

	Route{
		"CouchIdGet",
		strings.ToUpper("Get"),
		"/dev/couch/{Id}",
		CouchIdGet,
	},

	Route{
		"CouchIdPut",
		strings.ToUpper("Put"),
		"/dev/couch/{Id}",
		CouchIdPut,
	},

	Route{
		"CouchPost",
		strings.ToUpper("Post"),
		"/dev/couch",
		CouchPost,
	},

	Route{
		"StatusGet",
		strings.ToUpper("Get"),
		"/dev/status",
		StatusGet,
	},
}
