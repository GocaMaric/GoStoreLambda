package main

import (
	"context"
	"example.com/GoStoreLambda/awsgo"
	"example.com/GoStoreLambda/bd"
	"example.com/GoStoreLambda/handlers"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"os"
	"strings"
)

func main() {
	lambda.Start(ExecuteLambda)

}

func ExecuteLambda(ctx context.Context, request events.APIGatewayV2HTTPRequest) (*events.APIGatewayProxyResponse, error) {
	awsgo.InitializeAWS()

	if !ValidateParameters() {
		panic("Error with parameters. Need to get 'SecretName', 'UserPoolId', 'Region', 'UrlPrefix'")
	}

	var res *events.APIGatewayProxyResponse
	path := strings.Replace(request.RawPath, os.Getenv("UrlPrefix"), "", -1)
	method := request.RequestContext.HTTP.Method
	body := request.Body
	header := request.Headers

	bd.ReadSecret()

	status, message := handlers.Handlers(path, method, body, header, request)

	headersResp := map[string]string{
		"Content-Type": "application/json",
	}
	res = &events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       string(message),
		Headers:    headersResp,
	}
	return res, nil

}

func ValidateParameters() bool {
	_, bringsParameter := os.LookupEnv("SecretName")
	if !bringsParameter {
		return bringsParameter
	}
	_, bringsParameter = os.LookupEnv("UserPoolId")
	if !bringsParameter {
		return bringsParameter
	}
	_, bringsParameter = os.LookupEnv("Region")
	if !bringsParameter {
		return bringsParameter
	}
	_, bringsParameter = os.LookupEnv("UrlPrefix")
	if !bringsParameter {
		return bringsParameter
	}
	return bringsParameter

}
