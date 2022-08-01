package files

import "context"

type FileInterface interface {
	GetFiles(ctx context.Context, path string) ([]*S3File, error)
	PutFile(ctx context.Context, name string) (*string, error)
}
