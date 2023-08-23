package services

import (
	"context"
	"time"

	"github.com/Coderx44/s3_store/storage"
)

type Service interface {
	GetPreSignedPutUrl(ctx context.Context, objectName string, expiry time.Duration) (string, error)
	GetPreSignedGetUrl(ctx context.Context, objectName string, expiry time.Duration) (string, error)
}

type service struct {
	client storage.Storage
}

func NewStorageService(client storage.Storage) Service {
	return &service{client: client}
}

func (s *service) GetPreSignedPutUrl(ctx context.Context, objectName string, expiry time.Duration) (string, error) {
	presignedPutUrl, err := s.client.GetPreSignedPutUrl(ctx, objectName, expiry)
	if err != nil {
		return "", err
	}
	return presignedPutUrl, nil
}

func (s *service) GetPreSignedGetUrl(ctx context.Context, objectName string, expiry time.Duration) (string, error) {
	return "", nil
}
