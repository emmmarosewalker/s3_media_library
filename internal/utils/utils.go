package utils

import "github.com/aws/aws-lambda-go/events"

func CorsSuccessResponse(body string) *events.APIGatewayProxyResponse {
	return &events.APIGatewayProxyResponse{
		Body:       body,
		StatusCode: 200,
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		},
		IsBase64Encoded: true,
	}
}
