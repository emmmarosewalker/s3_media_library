package usecase

import (
	"context"
	"encoding/json"
	"s3_media_library/domain/files"
)

type FileUsecase struct {
	File files.FileInterface
}

func (u *FileUsecase) Get(ctx context.Context, path string) ([]*files.S3File, error) {
	res, err := u.File.GetFiles(ctx, path)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *FileUsecase) Put(ctx context.Context, name string) (*[]byte, error) {
	rawUrl, err := u.File.PutFile(ctx, name)

	if err != nil {
		return nil, err
	}

	url := files.S3PutUrl{
		Url: *rawUrl,
	}

	json, err := json.Marshal(url)

	if err != nil {
		return nil, err
	}

	return &json, nil
}
