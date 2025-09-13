package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	fullPath, pathErr := filepath.Abs(path)
	if pathErr != nil {
		return "", fmt.Errorf("error: %w", pathErr)
	}

	finalSize, err := getSize(fullPath, recursive, all)
	if err != nil {
		return "", fmt.Errorf("error: %w", err)
	}

	formattedSize := formatSize(finalSize, human)
	return formattedSize, nil
}

func formatSize(size int64, human bool) string {
	if !human || size < 1024 {
		return fmt.Sprintf("%dB", size)
	}

	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	unitIdx := 0
	finalSize := float64(size)
	for finalSize >= 1024 && unitIdx < len(units) {
		finalSize /= 1024
		unitIdx += 1
	}
	formattedSize := fmt.Sprintf("%.1f%s", finalSize, units[unitIdx])
	return formattedSize
}

func getSize(path string, recursive, all bool) (int64, error) {
	dataInfo, err := os.Lstat(path)
	if err != nil {
		return 0, fmt.Errorf("failed to get file/directory info: no such file or directory")
	}

	if !all && strings.HasPrefix(dataInfo.Name(), ".") {
		return 0, nil
	}

	if !dataInfo.IsDir() {
		return dataInfo.Size(), nil
	}

	var totalSize int64
	entries, dirErr := os.ReadDir(path)
	if dirErr != nil {
		return 0, fmt.Errorf("can't read directory: %w", dirErr)
	}
	for _, entry := range entries {
		if !recursive && entry.IsDir() {
			continue
		}
		size, err := getSize(filepath.Join(path, entry.Name()), recursive, all)
		if err != nil {
			return 0, fmt.Errorf("error: %w", err)
		}
		totalSize += size
	}
	return totalSize, nil
}
