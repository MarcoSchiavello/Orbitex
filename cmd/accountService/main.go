package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/marcoschiavello/Orbitex/internal/routers"
	T "github.com/marcoschiavello/Orbitex/internal/types"
)

func handler(req T.Request) (T.Response, error) {

	return routers.AccountRouter(req)
}

func main() {
	lambda.Start(handler)
}
