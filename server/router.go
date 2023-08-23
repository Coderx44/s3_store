package server

import (
	"net/http"

	"github.com/Coderx44/s3_store/services"
	"github.com/gorilla/mux"
)

func InitRouter(deps Dependencies) (router *mux.Router) {
	router = mux.NewRouter()

	router.HandleFunc("/get-signed-url", services.GetPreSignedPutUrl(deps.StoreService)).Methods(http.MethodGet)
	return

}
