package pathsize

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var units []string = []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	fullPath, pathErr := filepath.Abs(path)
	if pathErr != nil {
		return "", fmt.Errorf("error")
	}

	var res string
	finalSize, err := GetSize(fullPath, recursive, all)
	if err != nil {
		return "", fmt.Errorf("error: %w", err)
	}

	formattedSize := FormatSize(finalSize, human)
	res = fmt.Sprintf("%s\t%s", formattedSize, path)
	return res, nil
}

func FormatSize(size int64, human bool) string {
	if !human {
		return fmt.Sprintf("%dB", size)
	}
	unitIdx := 0
	var finalSize float64 = float64(size)
	for finalSize >= 1024 && unitIdx < len(units) {
		finalSize /= 1024
		unitIdx += 1
	}
	formattedSize := fmt.Sprintf("%.1f%s", finalSize, units[unitIdx])
	return formattedSize
}

func GetSize(path string, recursive, all bool) (int64, error) {
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
	entries, _ := os.ReadDir(path)
	for _, entry := range entries {
		if !recursive && entry.IsDir() {
			continue
		}
		size, err := GetSize(filepath.Join(path, entry.Name()), recursive, all)
		if err != nil {
			return 0, fmt.Errorf("error: %w", err)
		}
		totalSize += size
	}
	return totalSize, nil
}
