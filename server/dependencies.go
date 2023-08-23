package server

import (
	"github.com/Coderx44/s3_store/services"
	"github.com/Coderx44/s3_store/storage"
)

type Dependencies struct {
	StoreService services.Service
}

func initDependencies(client storage.Storage) Dependencies {
	storeService := services.NewStorageService(client)

	dep := Dependencies{
		StoreService: storeService,
	}
	return dep
}
