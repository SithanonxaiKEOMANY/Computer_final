package trails

import (
	"fmt"
	"os"
)

func DeleteImageFile(imagePath string) error {
	// Use os.Remove to delete the file
	err := os.Remove(imagePath)
	if err != nil {
		return fmt.Errorf("failed to delete image file: %v", err)
	}
	return nil
}
