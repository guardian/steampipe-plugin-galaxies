package store

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3 struct {
	bucket string
	client *s3.Client
}

func New(client *s3.Client, bucket string) S3 {
	return S3{
		client: client,
		bucket: bucket,
	}
}

func (store S3) Get(key string) ([]byte, error) {
	res, err := store.client.GetObject(context.Background(), &s3.GetObjectInput{
		Bucket: &store.bucket,
		Key:    &key,
	})

	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return data, nil
}
