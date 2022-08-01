package main

import (
	"context"
	"s3_media_library/application/usecase"
	"s3_media_library/infrastructure/persistance"
	"s3_media_library/internal/clients"
	"s3_media_library/internal/utils"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, event events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	name := event.QueryStringParameters["name"]
	filePersistance := persistance.NewFileRepository(clients.S3())

	fileUsecase := usecase.FileUsecase{
		File: filePersistance,
	}

	json, err := fileUsecase.Put(ctx, name)

	if err != nil {
		return nil, err
	}

	return utils.CorsSuccessResponse(string(*json)), nil
}

func main() {
	lambda.Start(HandleRequest)
}
