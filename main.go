package main

import (
	"github.com/Coderx44/s3_store/server"
)

// // MinIOCredentials holds the MinIO credentials
// type MinIOCredentials struct {
// 	Endpoint        string
// 	AccessKeyID     string
// 	SecretAccessKey string
// }

// // MinIOAPI is the API wrapper for MinIO
// type MinIOAPI struct {
// 	Client       *minio.Client
// 	BucketName   string
// 	SignedURLExp time.Duration
// }

// // NewMinIOAPI creates a new MinIOAPI instance
// func NewMinIOAPI(credential MinIOCredentials, bucketName string, signedURLExp time.Duration) (*MinIOAPI, error) {
// 	minioClient, err := minio.New(credential.Endpoint, &minio.Options{
// 		Creds:  credentials.NewStaticV4(credential.AccessKeyID, credential.SecretAccessKey, ""),
// 		Secure: false, // Change to true if using HTTPS
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &MinIOAPI{
// 		Client:       minioClient,
// 		BucketName:   bucketName,
// 		SignedURLExp: signedURLExp,
// 	}, nil
// }

// GetSignedURL generates a signed URL for uploading a file

// func (api *MinIOAPI) GetSignedURL() (string, error) {
// 	// Create a policy that allows any user to upload objects
// 	// policy := `{"Version": "2012-10-17","Statement": [{"Action": ["s3:PutObject"],"Effect": "Allow","Principal": {"AWS": ["*"]},"Resource": ["arn:aws:s3:::` + api.BucketName + `/*"],"Condition": {"StringEquals": {"s3:x-amz-acl": "private"}}}]}`
// 	// Initialize policy condition config.
// 	policy := minio.NewPostPolicy()

// 	// Apply upload policy restrictions:
// 	policy.SetBucket("newbucket")
// 	policy.SetKey("test")
// 	policy.SetExpires(time.Now().UTC().AddDate(0, 0, 10)) // expires in 10 days

// 	presignedURL, data, err := api.Client.PresignedPostPolicy(context.Background(), policy)
// 	if err != nil {
// 		log.Print(err.Error())
// 		return "", err
// 	}
// 	fmt.Print(data)
// 	return presignedURL.String(), nil
// }

// APIHandler is the HTTP handler for getting a signed URL
// func (api *MinIOAPI) APIHandler(w http.ResponseWriter, r *http.Request) {
// objectName := "object.pdf" // Adjust the object name as needed

// 	signedURL, err := api.GetSignedURL()
// 	if err != nil {
// 		http.Error(w, "Error generating signed URL", http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "text/plain")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte(signedURL))
// }

func main() {
	// credentials := MinIOCredentials{
	// 	Endpoint:        "192.168.1.117:9000",
	// 	AccessKeyID:     "JSw7V41myXOrgWHr5MHl",
	// 	SecretAccessKey: "dlqxBh8cMDIx8R3Lzye3DSjEgC9CRBtgxDqYjkrY",
	// }
	// bucketName := "newbucket"
	// signedURLExp := 1 * time.Hour // Set the expiration time for the signed URL

	// minioAPI, err := NewMinIOAPI(credentials, bucketName, signedURLExp)
	// if err != nil {
	// 	fmt.Println("Error initializing MinIO API:", err)
	// 	return
	// }

	// router := mux.NewRouter()
	// router.HandleFunc("/get-signed-url", minioAPI.APIHandler)

	// fmt.Println("Server started on :8080")
	// http.Handle("/", router)
	// http.ListenAndServe(":8080", nil)
	err := server.StartApp()
	if err != nil {
		panic(err)
	}
}
