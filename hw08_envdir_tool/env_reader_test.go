package main

import (
	"io/fs"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	t.Run("Non-existing file", func(t *testing.T) {
		_, err := ReadDir("not_exists-Env.files_dir")
		require.ErrorIs(t, err, fs.ErrNotExist)
	})

	t.Run("Empty dir", func(t *testing.T) {
		testDir := "testdata/emtdir"

		err := os.Mkdir(testDir, os.ModePerm)
		defer func() {
			if err := os.Remove(testDir); err != nil {
				log.Fatal(err)
			}
		}()
		if err != nil {
			log.Fatal(err)
		}
		env, err := ReadDir(testDir)
		require.Nil(t, err)
		require.Empty(t, env)
	})

	t.Run("Success read", func(t *testing.T) {
		env := Environment{
			"BAR": {
				Value:      "bar",
				NeedRemove: false,
			},
			"EMPTY": {
				Value:      "",
				NeedRemove: false,
			},
			"FOO": {
				Value:      "   foo\nwith new line",
				NeedRemove: false,
			},
			"HELLO": {
				Value:      `"hello"`,
				NeedRemove: false,
			},
			"UNSET": {
				Value:      "",
				NeedRemove: true,
			},
		}
		rEnv, err := ReadDir("testdata/env")
		require.Nil(t, err)
		require.Equal(t, env, rEnv)
	})
}
