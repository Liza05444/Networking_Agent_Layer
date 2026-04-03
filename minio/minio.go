package minio

import (
	"context"
	"io"

	"agent/config"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var client *minio.Client

func Init() error {
	var err error

	client, err = minio.New(config.MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.MinioAccessKey, config.MinioSecretKey, ""),
		Secure: false,
	})

	return err
}

func DownloadImage(objectName string) ([]byte, error) {
	ctx := context.Background()

	obj, err := client.GetObject(ctx, config.BucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	defer obj.Close()

	return io.ReadAll(obj)
}
