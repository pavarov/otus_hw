package main

import (
	"flag"
	"log"
	"os"
)

var (
	from, to      string
	limit, offset int64
)

func init() {
	flag.StringVar(&from, "from", "", "file to read from")
	flag.StringVar(&to, "to", "", "file to write to")
	flag.Int64Var(&limit, "limit", 0, "limit of bytes to copy")
	flag.Int64Var(&offset, "offset", 0, "offset in input file")
}

func inputValidate() error {
	if from == "" || to == "" {
		return ErrUnsupportedFile
	}
	s, err := os.Stat(from)
	if err != nil {
		return err
	}
	fileSize := s.Size()
	if fileSize == 0 {
		return ErrFileWithUnknownSize
	}
	if offset > fileSize {
		return ErrOffsetExceedsFileSize
	}
	return nil
}

func main() {
	flag.Parse()
	if err := inputValidate(); err != nil {
		log.Fatal(err)
	}
	if errCopy := Copy(from, to, offset, limit); errCopy != nil {
		log.Fatal(errCopy)
	}
}
