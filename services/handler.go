package services

import (
	"errors"
	"io"
	"net/http"

	"github.com/gabriel-vasile/mimetype"
)

func ValidateUpload(file io.Reader) error {
	detectedMIME, err := mimetype.DetectReader(file)
	if err != nil {
		return err
	}

	allowedTypes := []string{"image/jpeg", "image/png", "application/pdf"} // Add more allowed types as needed

	valid := false
	for _, allowedType := range allowedTypes {
		if detectedMIME.String() == allowedType {
			valid = true
			break
		}
	}

	if !valid {
		return errors.New("invalid file type")
	}

	return nil
}

func GetPreSignedPutUrl(service Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseMultipartForm(20 << 20)
		if err != nil {
			http.Error(w, "Unable to parse form data", http.StatusBadRequest)
			return
		}
		// Get the uploaded file
		file, handler, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Unable to retrieve file from form", http.StatusBadRequest)
			return
		}
		defer file.Close()
		// Validate the uploaded file's MIME type
		err = ValidateUpload(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		presignedPutUrl, err := service.GetPreSignedPutUrl(r.Context(), handler.Filename, 0)
		if err != nil {
			http.Error(w, "Error generating signed URL", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(presignedPutUrl))
	}
}
