package pathsize

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var units []string = []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	fullPath, _ := filepath.Abs(path)
	dataInfo, err := os.Lstat(fullPath)
	if err != nil {
		return "", fmt.Errorf("failed to get file info: no such file or directory")
	}

	var res string
	if !dataInfo.IsDir() {
		fileSize := FormatSize(dataInfo.Size(), human)
		res = fmt.Sprintf("%s\t%s", fileSize, dataInfo.Name())
		return res, nil
	}

	entries, dirErr := os.ReadDir(fullPath)
	if dirErr != nil {
		return "", fmt.Errorf("failed to get directory info: %w", dirErr)
	}

	var filesSize int64
	if all {
		for _, entry := range entries {
			dataInfo, err := entry.Info()
			if err == nil && !dataInfo.IsDir() {
				filesSize += dataInfo.Size()
			}
		}
	} else {
		for _, entry := range entries {
			dataInfo, err := entry.Info()
			if err == nil && !dataInfo.IsDir() {
				name := dataInfo.Name()
				if !strings.HasPrefix(name, ".") {
					filesSize += dataInfo.Size()
				}
			}
		}
	}
	formattedSize := FormatSize(filesSize, human)
	res = fmt.Sprintf("%s\t%s", formattedSize, dataInfo.Name())
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
