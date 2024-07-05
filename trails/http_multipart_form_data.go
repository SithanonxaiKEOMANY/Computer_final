package trails

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"io/ioutil"
)

// MaximumFileSize defines the maximum allowed file size in bytes (10MB)
const MaximumFileSize = 10 * 1024 * 1024 // 10MB

// HandleMultipartFormData extracts image data from multipart form data
func HandleMultipartFormData(ctx *fiber.Ctx) ([]byte, error) {
	// Parse the multipart form data
	form, err := ctx.MultipartForm()
	if err != nil {
		return nil, err
	}

	// Retrieve the file from the form data
	files := form.File["image"]
	if len(files) == 0 {
		return nil, errors.New("no image file uploaded")
	}

	// Get the first file from the slice
	file := files[0]

	// Check if file size exceeds the maximum allowed size
	if file.Size > MaximumFileSize {
		return nil, errors.New("file size exceeds the maximum allowed size")
	}

	// Open the file from the form
	uploadedFile, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer uploadedFile.Close()

	// Read the file data into a byte slice
	imageData, err := ioutil.ReadAll(uploadedFile)
	if err != nil {
		return nil, err
	}

	return imageData, nil
}
