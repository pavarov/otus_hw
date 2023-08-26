package main

import (
	"bufio"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

func closeFileDesc(f *os.File) {
	if errClose := f.Close(); errClose != nil {
		log.Fatal(errClose)
	}
}

func parseEnvFromFile(f *os.File) string {
	var v string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		v = scanner.Text()
		break
	}

	v = strings.TrimRight(v, "\t ")
	v = strings.ReplaceAll(v, "\x00", "\n")
	return v
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	envs := make(Environment)
	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			fileEnvName := info.Name()

			if strings.Contains(fileEnvName, "=") {
				return nil
			}

			envVal := new(EnvValue)
			if info.Size() == 0 {
				envVal.NeedRemove = true
				envs[fileEnvName] = *envVal
				return nil
			}

			f, err := os.Open(path)
			defer closeFileDesc(f)
			if err != nil {
				return err
			}

			envVal.Value = parseEnvFromFile(f)
			envVal.NeedRemove = false

			envs[fileEnvName] = *envVal
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	// Place your code here
	return envs, nil
}
