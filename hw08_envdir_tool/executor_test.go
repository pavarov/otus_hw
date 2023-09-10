package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	t.Run("Success remove var", func(t *testing.T) {
		envName := "OLD_ENV_VARIABLE"
		envVal := "some val"
		require.Nil(t, os.Setenv(envName, envVal))
		v, ok := os.LookupEnv(envName)
		require.True(t, ok)
		require.Equal(t, envVal, v)

		env := Environment{
			envName: {
				Value:      envVal,
				NeedRemove: true,
			},
		}
		require.Nil(t, fillEnv(env))
		_, ok = os.LookupEnv(envName)
		require.False(t, ok)
	})

	t.Run("Success set var", func(t *testing.T) {
		envNewName := "NEW_ENV_VAR"
		envVal := "some val"
		_, ok := os.LookupEnv(envNewName)
		require.False(t, ok)

		env := Environment{
			envNewName: {
				Value:      envVal,
				NeedRemove: false,
			},
		}
		require.Nil(t, fillEnv(env))

		v, ok := os.LookupEnv(envNewName)
		require.True(t, ok)
		require.Equal(t, v, envVal)
	})

	t.Run("success", func(t *testing.T) {
		code := RunCmd([]string{"ls"}, nil)
		require.Equal(t, 0, code)
	})

	t.Run("Executable not found", func(t *testing.T) {
		code := RunCmd([]string{"ls_lah"}, nil)
		require.Equal(t, 1, code)
	})

	t.Run("Command not found", func(t *testing.T) {
		code := RunCmd([]string{"/bin/bash", "tmm"}, nil)
		require.Equal(t, 127, code)
	})
}
