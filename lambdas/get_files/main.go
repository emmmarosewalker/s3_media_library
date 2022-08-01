package main

import (
	"context"
	"encoding/json"
	"s3_media_library/application/usecase"
	"s3_media_library/infrastructure/persistance"
	"s3_media_library/internal/clients"
	"s3_media_library/internal/utils"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, event events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	path := event.QueryStringParameters["path"]
	filePersistance := persistance.NewFileRepository(clients.S3())

	fileUsecase := usecase.FileUsecase{
		File: filePersistance,
	}

	files, err := fileUsecase.Get(ctx, path)

	if err != nil {
		return nil, err
	}

	res, err := json.Marshal(files)

	if err != nil {
		return nil, err
	}

	return utils.CorsSuccessResponse(string(res)), nil
}

func main() {
	lambda.Start(HandleRequest)
}
