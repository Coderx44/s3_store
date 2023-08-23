package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Coderx44/s3_store/storage"
)

func StartApp() (err error) {
	credentials := storage.MinIOCredentials{
		Endpoint:        "192.168.1.117:9000",
		AccessKeyID:     "JSw7V41myXOrgWHr5MHl",
		SecretAccessKey: "dlqxBh8cMDIx8R3Lzye3DSjEgC9CRBtgxDqYjkrY",
	}
	bucketName := "newbucket"
	signedURLExp := 1 * time.Hour // Set the expiration time for the signed URL

	client, err := storage.NewMinIOAPI(credentials, bucketName, signedURLExp)
	if err != nil {
		log.Printf("error creating client: %v", err)
		return
	}
	deps := initDependencies(client)
	router := InitRouter(deps)
	fmt.Println("Server started on :8080")
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
	return
}
