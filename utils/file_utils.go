package utils

import "path/filepath"

func GetTypeFromFilename(filename string) (string, error) {
	ext := filepath.Ext(filename)

	return ext, nil
}
