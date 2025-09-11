package tests

import (
	"code"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetPathSizeOfFile(t *testing.T) {
	fileSize, err := code.GetPathSize("../testdata/file1.txt", false, false, false)
	require.NoError(t, err, "File exists, but error was returned")
	require.Equal(t, "13B", fileSize, "Wrong size of file")
}

func TestGetPathSizeOfDirectory(t *testing.T) {
	directorySize, err := code.GetPathSize("../testdata/directory1", false, false, false)
	require.NoError(t, err, "Directory exists, but error was returned")
	require.Equal(t, "39B", directorySize, "Wrong directory size")
}

func TestGetPatSizeOfEmptyDirectory(t *testing.T) {
	emptyDirSize, err := code.GetPathSize("../testdata/emptydir", false, false, false)
	require.NoError(t, err, "Directory exists, but error was returned")
	require.Equal(t, "0B", emptyDirSize, "Size is not 0")
}

func TestGetPathSizeOfNotExistsFile(t *testing.T) {
	_, err := code.GetPathSize("../testdata/random", false, false, false)
	require.EqualError(t, err, "error: failed to get file/directory info: no such file or directory")
}

func TestGetDirSizeWithoutHiddenFiles(t *testing.T) {
	dirSize, _ := code.GetPathSize("../testdata/directory1", false, false, false)
	require.Equal(t, "39B", dirSize, "size not without hidden files")
}

func TestGetDirSizeWithHiddenFiles(t *testing.T) {
	dirSize, _ := code.GetPathSize("../testdata/directory1", false, false, true)
	require.Equal(t, "52B", dirSize, "size not without hidden files")
}

func TestGetDirSizeWithRecursiveAndHidden(t *testing.T) {
	dirSize, _ := code.GetPathSize("../testdata", true, false, true)
	require.Equal(t, "6213B", dirSize, "wrong size with recursive and hidden flags")
}

func TestGetDirSizeWithoutRecursiveAndHidden(t *testing.T) {
	dirSize, _ := code.GetPathSize("../testdata/directory1", false, false, false)
	require.Equal(t, "39B", dirSize, "wrong size with recursive and hidden flags")
}

func TestGetDirSizeWithoutRecursiveAndWithHidden(t *testing.T) {
	dirSize, _ := code.GetPathSize("../testdata/directory1", false, false, true)
	require.Equal(t, "52B", dirSize, "wrong size with recursive and hidden flags")
}
