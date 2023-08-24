package services

import (
	"net/http"
)

func GetPreSignedPutUrl(service Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queryParams := r.URL.Query()
		objectName := queryParams.Get("object_name")
		// Validate if the object_name parameter is empty
		if objectName == "" {
			http.Error(w, "object_name parameter is required", http.StatusBadRequest)
			return
		}
		presignedPutUrl, err := service.GetPreSignedPutUrl(r.Context(), objectName)
		if err != nil {
			http.Error(w, "Error generating signed URL", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(presignedPutUrl))
	}
}

func GetPreSignedGetUrl(service Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queryParams := r.URL.Query()
		objectName := queryParams.Get("object_name")
		// Validate if the object_name parameter is empty
		if objectName == "" {
			http.Error(w, "object_name parameter is required", http.StatusBadRequest)
			return
		}
		presignedGetUrl, err := service.GetPreSignedGetUrl(r.Context(), objectName)
		if err != nil {
			http.Error(w, "Error generating signed URL", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(presignedGetUrl))
	}
}
