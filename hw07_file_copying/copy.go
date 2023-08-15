package main

import (
	"errors"
	"io"
	"log"
	"os"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrFileWithUnknownSize   = errors.New("file with unknown size")
)

func closeFileDesc(f *os.File) {
	if errClose := f.Close(); errClose != nil {
		log.Fatal(errClose)
	}
}

func copyFile(out, f *os.File, limit int64) error {
	if _, err := io.CopyN(out, f, limit); err != nil {
		if errors.Is(err, io.EOF) {
			return nil
		}
		return err
	}
	return nil
}

func Copy(fromPath, toPath string, offset, limit int64) error {
	f, err := os.Open(fromPath)
	defer closeFileDesc(f)
	if err != nil {
		return err
	}

	if _, errS := f.Seek(offset, io.SeekStart); errS != nil {
		return err
	}

	out, err := os.Create(toPath)
	defer closeFileDesc(out)
	if err != nil {
		return err
	}

	if limit == 0 {
		s, _ := f.Stat()
		limit = s.Size()
	}
	return copyFile(out, f, limit)
}
