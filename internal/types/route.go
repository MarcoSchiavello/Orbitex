package types

type RouteHandler func(Request) (Response, error)
