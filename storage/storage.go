package storage

import (
	"context"
	"time"
)

type Storage interface {
	GetPreSignedPutUrl(ctx context.Context, objectName string, expiry time.Duration) (string, error)
	GetPreSignedGetUrl(ctx context.Context, objectName string, expiry time.Duration) (string, error)
}
