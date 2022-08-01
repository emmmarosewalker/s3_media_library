package persistance

import (
	"context"
	"log"
	"s3_media_library/domain/files"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type fileRepositoryImp struct {
	s3Conn *s3.S3
}

func NewFileRepository(s3Conn *s3.S3) files.FileInterface {
	return &fileRepositoryImp{
		s3Conn,
	}
}

func (f *fileRepositoryImp) GetFiles(ctx context.Context, path string) ([]*files.S3File, error) {
	s3Input := s3.ListObjectsV2Input{
		Bucket:    aws.String(("emmy-codes-media-library")),
		Delimiter: aws.String("/"),
		Prefix:    aws.String(path),
	}

	res, err := f.s3Conn.ListObjectsV2WithContext(ctx, &s3Input)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	contents := res.Contents

	result := []*files.S3File{}

	commonPrefixes := res.CommonPrefixes

	if len(commonPrefixes) == 0 {
		log.Print("no common prefixes")
	}

	for _, prefix := range commonPrefixes {
		log.Print(*prefix)
		result = append(result, &files.S3File{
			Name:         *prefix.Prefix,
			LastModified: "",
			Size:         "",
			Type:         "folder",
		})
	}

	for _, s := range contents {
		result = append(result, &files.S3File{
			Name:         *s.Key,
			LastModified: s.LastModified.String(),
			Size:         strconv.Itoa(int(*s.Size)),
			Type:         "file",
		})
	}

	return result, nil
}

func (f *fileRepositoryImp) PutFile(ctx context.Context, name string) (*string, error) {
	s3Input := s3.PutObjectInput{
		Bucket: aws.String("emmy-codes-media-library"),
		Key:    aws.String(name),
	}

	req, _ := f.s3Conn.PutObjectRequest(&s3Input)

	url, _, err := req.PresignRequest(time.Minute * 15)

	if err != nil {
		return nil, err
	}

	return &url, nil
}
