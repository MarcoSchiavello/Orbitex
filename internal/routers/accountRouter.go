package routers

import (
	"strings"

	T "github.com/marcoschiavello/Orbitex/internal/types"
)

func AccountRouter(req T.Request) (T.Response, error) {
	routes := map[string]T.RouteHandler{
		//------------------ Account-Routes ---------------------
		"GET:/users": func(req T.Request) (T.Response, error) {
			return T.Response{
				Body:       req.Path,
				StatusCode: 200,
			}, nil
		},

		"GET:/users/" + req.PathParameters["id"]: func(req T.Request) (T.Response, error) {
			return T.Response{
				Body:       req.Path + "    " + req.PathParameters["id"],
				StatusCode: 200,
			}, nil
		},

		"POST:/users": func(req T.Request) (T.Response, error) {
			return T.Response{
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
		return T.Response{
			StatusCode: 404,
		}, nil
	}
}
