package main

import (
	"Vizinhos_Back_End/Handler"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
)

func handleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch req.Path {
	case "/customer/":
		return Handler.GetCustomerDataHandler(req)
	case "/seller/":
		return Handler.GetSellerDataHandler(req)
	case "/register/user/":
		return Handler.RegisterUserHandler(req)
	default:
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusNotFound,
			Body:       "Not Found",
		}, nil
	}
}

func main() {
	lambda.Start(handleRequest)
}
