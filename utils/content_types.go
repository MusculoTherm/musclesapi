package utils

import (
	"errors"
	"fmt"
)

func CheckImageContentType(contentType string) (string, error) {
	switch contentType {
	case "image/bmp":
		return ".bmp", nil
	case "image/jpeg":
		return ".jpg", nil
	case "image/gif":
		return ".gif", nil
	case "image/png":
		return ".png", nil
	default:
		return "", errors.New(fmt.Sprintf("%s is not a supported content type.", contentType))
	}
}
