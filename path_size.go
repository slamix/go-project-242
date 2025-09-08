package pathsize

import (
	"fmt"
	"os"
)

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	fileInfo, err := os.Lstat(path)
	if err != nil {
		return "", fmt.Errorf("failed to get file info: %w", err)
	}

	if !fileInfo.IsDir() {
		res := fmt.Sprintf("%dB\t%s", fileInfo.Size(), fileInfo.Name())
		return res, nil
	}

	entries, dirErr := os.ReadDir(path)
	if dirErr != nil {
		return "", fmt.Errorf("failed to get directory info: %w", dirErr)
	}

	var filesSize int64
	for _, entry := range entries {
		dataInfo, err := entry.Info()
		if err == nil || !dataInfo.IsDir() {
			filesSize += dataInfo.Size()
		}
	}

	res := fmt.Sprintf("%dB\t%s", filesSize, path)
	return res, nil
}
