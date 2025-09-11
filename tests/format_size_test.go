package tests

import (
	"code"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFormatSizeWithoutHuman(t *testing.T) {
	res := code.FormatSize(4939042, false)
	require.Equal(t, res, "4939042B")
}

func TestFormatSizeWithHuman(t *testing.T) {
	res := code.FormatSize(4939042, true)
	require.Equal(t, res, "4.7MB")
}

func TestFormatSizeOnBorderValue(t *testing.T) {
	res := code.FormatSize(1048576, true)
	require.Equal(t, res, "1.0MB")
}
