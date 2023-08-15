package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	t.Run("without file args", func(t *testing.T) {
		require.ErrorIs(t, inputValidate(), ErrUnsupportedFile)
	})

	t.Run("src file not exists", func(t *testing.T) {
		from = "f.txt"
		to = "t.txt"
		require.ErrorIs(t, inputValidate(), os.ErrNotExist)
	})

	t.Run("src file unknown size", func(t *testing.T) {
		from = "/dev/null"
		to = "s.txt"
		require.ErrorIs(t, inputValidate(), ErrFileWithUnknownSize)
	})

	t.Run("offset too large", func(t *testing.T) {
		from = "testdata/out_offset0_limit10.txt"
		to = "t1.txt"
		offset = 11
		require.ErrorIs(t, inputValidate(), ErrOffsetExceedsFileSize)
	})

	t.Run("right args", func(t *testing.T) {
		from = "testdata/out_offset0_limit10.txt"
		to = "t2.txt"
		offset = 0
		limit = 1
		require.Nil(t, inputValidate())
	})
}
