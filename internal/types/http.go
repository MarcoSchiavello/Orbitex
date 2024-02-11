package types

import (
	"github.com/aws/aws-lambda-go/events"
)

type Request events.APIGatewayProxyRequest

type Response events.APIGatewayProxyResponse

type ReqContext events.APIGatewayProxyRequestContext

type ReqIdentity events.APIGatewayRequestIdentity
