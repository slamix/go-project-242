package tests

import (
	code "code"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetPathSizeOfFile(t *testing.T) {
	fileSize, err := code.GetPathSize("../testdata/file1.txt", false, false, false)
	require.NoError(t, err, "File exists, but error was returned")
	require.Equal(t, "13B\tfile1.txt", fileSize, "Wrong size of file")
}

func TestGetPathSizeOfDirectory(t *testing.T) {
	directorySize, err := code.GetPathSize("../testdata/directory1", false, false, false)
	require.NoError(t, err, "Directory exists, but error was returned")
	require.Equal(t, "39B\tdirectory1", directorySize, "Wrong directory size")
}

func TestGetPatSizeOfEmptyDirectory(t *testing.T) {
	emptyDirSize, err := code.GetPathSize("../testdata/emptydir", false, false, false)
	require.NoError(t, err, "Directory exists, but error was returned")
	require.Equal(t, "0B\temptydir", emptyDirSize, "Size is not 0")
}

func TestGetPathSizeOfNotExistsFile(t *testing.T) {
	_, err := code.GetPathSize("../testdata/random", false, false, false)
	require.EqualError(t, err, "failed to get file info: no such file or directory")
}

func TestGetDirSizeWithoutHiddenFiles(t *testing.T) {
	dirSize, _ := code.GetPathSize("../testdata/directory1", false, false, false)
	require.Equal(t, "39B\tdirectory1", dirSize, "size not without hidden files")
}

func TestGetDirSizeWithHiddenFiles(t *testing.T) {
	dirSize, _ := code.GetPathSize("../testdata/directory1", false, false, true)
	require.Equal(t, "52B\tdirectory1", dirSize, "size not without hidden files")
}
