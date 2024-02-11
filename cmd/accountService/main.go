package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	r "github.com/marcoschiavello/Orbitex/internal/routers"
	t "github.com/marcoschiavello/Orbitex/internal/types"
)

func handler(req t.Request) (t.Response, error) {

	return r.AccountRouter(req)
}

func main() {
	lambda.Start(handler)
}
