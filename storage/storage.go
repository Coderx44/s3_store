package storage

import (
	"context"
)

type Storage interface {
	GetPreSignedPutUrl(ctx context.Context, objectName string) (string, error)
	GetPreSignedGetUrl(ctx context.Context, objectName string) (string, error)
}
