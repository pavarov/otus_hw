package hw10programoptimization

import (
	"archive/zip"
	"testing"

	"github.com/stretchr/testify/require"
)

func BenchmarkGetDomainStat(b *testing.B) {
	b.StopTimer()

	r, err := zip.OpenReader("testdata/users.dat.zip")
	require.NoError(b, err)
	defer func(r *zip.ReadCloser) {
		err := r.Close()
		if err != nil {
			panic(err)
		}
	}(r)

	require.Equal(b, 1, len(r.File))

	data, err := r.File[0].Open()
	require.NoError(b, err)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = GetDomainStat(data, "biz")
	}
}
