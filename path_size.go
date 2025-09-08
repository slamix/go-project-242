package pathsize

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	fullPath, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("failed to get absolute path: %w", err)
	}
	dataInfo, err := os.Lstat(fullPath)
	if err != nil {
		return "", fmt.Errorf("failed to get file info: no such file or directory")
	}

	if !dataInfo.IsDir() {
		res := fmt.Sprintf("%dB\t%s", dataInfo.Size(), dataInfo.Name())
		return res, nil
	}

	entries, dirErr := os.ReadDir(fullPath)
	if dirErr != nil {
		return "", fmt.Errorf("failed to get directory info: %w", dirErr)
	}

	var filesSize int64
	for _, entry := range entries {
		dataInfo, err := entry.Info()
		if err == nil && !dataInfo.IsDir() {
			filesSize += dataInfo.Size()
		}
	}

	res := fmt.Sprintf("%dB\t%s", filesSize, dataInfo.Name())
	return res, nil
}
