package routes

import (
	"strings"

	t "github.com/marcoschiavello/Orbitex/internal/types"
)

func AccountRouter(req t.Request) (t.Response, error) {
	routes := map[string]t.RouteHandler{
		//------------------ Account-Routes ---------------------
		"GET:/users": func(req t.Request) (t.Response, error) {
			return t.Response{
				Body:       req.Path,
				StatusCode: 200,
			}, nil
		},

		"GET:/users/" + req.PathParameters["id"]: func(req t.Request) (t.Response, error) {
			return t.Response{
				Body:       req.Path + "    " + req.PathParameters["id"],
				StatusCode: 200,
			}, nil
		},

		"POST:/users": func(req t.Request) (t.Response, error) {
			return t.Response{
				Body:       req.Path,
				StatusCode: 200,
			}, nil
		},
	}

	//--------------------- Logic ----------------------

	key := strings.ToUpper(req.HTTPMethod) + ":" + strings.TrimSuffix(req.Path, "/")

	function, found := routes[key]

	if found {
		return function(req)
	} else {
		return t.Response{
			StatusCode: 404,
		}, nil
	}
}
