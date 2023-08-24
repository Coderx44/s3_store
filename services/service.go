package services

import (
	"context"

	"github.com/Coderx44/s3_store/storage"
)

type Service interface {
	GetPreSignedPutUrl(ctx context.Context, objectName string) (string, error)
	GetPreSignedGetUrl(ctx context.Context, objectName string) (string, error)
}

type service struct {
	client storage.Storage
}

func NewStorageService(client storage.Storage) Service {
	return &service{client: client}
}

func (s *service) GetPreSignedPutUrl(ctx context.Context, objectName string) (string, error) {

	presignedPutUrl, err := s.client.GetPreSignedPutUrl(ctx, objectName)
	if err != nil {
		return "", err
	}
	return presignedPutUrl, nil
}

func (s *service) GetPreSignedGetUrl(ctx context.Context, objectName string) (string, error) {

	presignedGetUrl, err := s.client.GetPreSignedGetUrl(ctx, objectName)
	if err != nil {
		return "", err
	}
	return presignedGetUrl, nil
}
