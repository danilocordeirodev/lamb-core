package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/danilocordeirodev/lamb-core/awsgo"
	"github.com/danilocordeirodev/lamb-core/db"
)

func main() {
	lambda.Start(LambdaExecution)
}

func LambdaExecution(ctx context.Context, request events.APIGatewayV2HTTPRequest) (*events.APIGatewayProxyResponse, error){
	awsgo.InitializeAWS()

	if !ValidateParameters() {
		fmt.Println("Error in parameters. Should send 'SecretManager', 'UserPoolId', 'Region', 'UrlPrefix'")
		err := errors.New("Error in parameters. Should send 'SecretManager', 'UserPoolId', 'Region', 'UrlPrefix'")
	}

	var res *events.APIGatewayProxyResponse
	prefix := os.Getenv("UrlPrefix")
	path := strings.Replace(request.RawPath, prefix, "", -1)
	method := request.RequestContext.HTTP.Method
	body := request.Body
	header := request.Headers
	
	db.ReadSecret()

	headersResp := map[string]string {
		"Content-Type": "application/json",
	}

	res = &events.APIGatewayProxyResponse{
		StatusCode: status,
		Body: string(message),
		Headers: headersResp,
	}

	return res, nil

}

func ValidateParameters() bool {
	_, checkParameter := os.LookupEnv("SecretName")
	if !checkParameter {
		return checkParameter
	}

	_, checkParameter = os.LookupEnv("UserPoolId")
	if !checkParameter {
		return checkParameter
	}

	_, checkParameter = os.LookupEnv("Region")
	if !checkParameter {
		return checkParameter
	}

	_, checkParameter = os.LookupEnv("UrlPrefix")
	if !checkParameter {
		return checkParameter
	}

	return checkParameter
}