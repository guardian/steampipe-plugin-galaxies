package store

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"io"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3 struct {
	bucket string
	prefix string
	client *s3.Client
}

func New(client *s3.Client, bucket, prefix string) S3 {
	return S3{
		client: client,
		bucket: bucket,
		prefix: prefix,
	}
}

func (store S3) Get(key string) ([]byte, error) {
	fullKey := store.prefix + "/" + key

	res, err := store.client.GetObject(context.Background(), &s3.GetObjectInput{
		Bucket: &store.bucket,
		Key:    &fullKey,
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

func GalaxiesS3(ctx context.Context) (S3, error) {
	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion("eu-west-1"),
		config.WithSharedConfigProfile("deployTools"),
	)

	if err != nil {
		return S3{}, fmt.Errorf("unable to load AWS config, %w", err)
	}

	client := s3.NewFromConfig(cfg)
	bucket := os.Getenv("GALAXIES_BUCKET")
	s := New(client, bucket, "galaxies.gutools.co.uk/data")
	return s, nil
}
