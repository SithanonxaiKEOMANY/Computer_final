package trails

//import (
//	"github.com/gofiber/fiber/v2"
//	"io/ioutil"
//	"mime/multipart"
//	"os"
//	"path/filepath"
//)
//
//// HandleMultipartFormData parses the multipart form data and saves the uploaded image locally.
//func HandleMultipartFormData(ctx *fiber.Ctx) ([]byte, error) {
//	// Get the file from the form data
//	file, _, err := ctx.FormFile("image")
//	if err != nil {
//		return nil, err
//	}
//	defer file.Close()
//
//	// Read the file data into a byte slice
//	imageData, err := ioutil.ReadAll(file)
//	if err != nil {
//		return nil, err
//	}
//
//	// Define the directory structure
//	directory := filepath.Join("assets", "ce_it", "2024")
//
//	// Create the directory if it doesn't exist
//	err = os.MkdirAll(directory, os.ModePerm)
//	if err != nil {
//		return nil, err
//	}
//
//	// Create a unique filename
//	filename := ctx.FormValue("student_id") + ".png" // You may want to generate a unique filename
//
//	// Create the file
//	path := filepath.Join(directory, filename)
//	err = ioutil.WriteFile(path, imageData, 0644)
//	if err != nil {
//		return nil, err
//	}
//
//	// Return the image data
//	return imageData, nil
//}
