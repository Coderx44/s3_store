package storage

import (
	"context"
	"log"
	"time"

	"github.com/minio/minio-go/v7"
	credentials "github.com/minio/minio-go/v7/pkg/credentials"
)

// MinIOCredentials holds the MinIO credentials
type MinIOCredentials struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
}

// MinIOAPI is the API wrapper for MinIO
type MinIOAPI struct {
	Client       *minio.Client
	BucketName   string
	SignedURLExp time.Duration
}

// NewMinIOAPI creates a new MinIOAPI instance
func NewMinIOAPI(credential MinIOCredentials, bucketName string, signedURLExp time.Duration) (Storage, error) {
	minioClient, err := minio.New(credential.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(credential.AccessKeyID, credential.SecretAccessKey, ""),
		Secure: false, // Change to true if using HTTPS
	})
	if err != nil {
		return nil, err
	}

	return &MinIOAPI{
		Client:       minioClient,
		BucketName:   bucketName,
		SignedURLExp: signedURLExp,
	}, nil
}

func (api *MinIOAPI) GetPreSignedPutUrl(ctx context.Context, objectName string) (string, error) {

	presignedURL, err := api.Client.PresignedPutObject(ctx, api.BucketName, objectName, api.SignedURLExp)
	if err != nil {
		log.Printf("Error generating presigned URL: %v", err)
		return "", err
	}
	return presignedURL.String(), nil
}

func (api *MinIOAPI) GetPreSignedGetUrl(ctx context.Context, objectName string) (string, error) {
	presignedUrl, err := api.Client.PresignedGetObject(ctx, api.BucketName, objectName, api.SignedURLExp, nil)
	if err != nil {
		log.Printf("Error generating presigned URL: %v", err)
		return "", err
	}
	return presignedUrl.String(), nil
}
