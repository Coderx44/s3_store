package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Coderx44/s3_store/bootstrap"
	"github.com/Coderx44/s3_store/storage"
)

func StartApp() (err error) {
	env := bootstrap.NewEnv()

	// credentials := storage.MinIOCredentials{
	// 	Endpoint:        "192.168.1.117:9000",
	// 	AccessKeyID:     "JSw7V41myXOrgWHr5MHl",
	// 	SecretAccessKey: "dlqxBh8cMDIx8R3Lzye3DSjEgC9CRBtgxDqYjkrY",
	// }

	credentials := storage.MinIOCredentials{
		Endpoint:        env.CloudEndpoint,
		AccessKeyID:     env.CloudAccessKey,
		SecretAccessKey: env.CloudSecretKey,
	}
	bucketName := env.CloudBucketName
	signedURLExp := time.Duration(env.SignedUrlExpiryMinute) * time.Minute

	client, err := storage.NewMinIOAPI(credentials, bucketName, signedURLExp)
	if err != nil {
		log.Printf("error creating client: %v", err)
		return
	}
	deps := initDependencies(client)
	router := InitRouter(deps)
	fmt.Printf("Server started on :%s", env.AppPort)
	http.Handle("/", router)
	http.ListenAndServe(fmt.Sprintf(":%s", env.AppPort), nil)
	return
}
