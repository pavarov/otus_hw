package main

import (
	"log"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		log.Fatal("go-envdir usage: min 2 arguments (env dir, command)")
	}

	envs, err := ReadDir(args[1])
	if err != nil {
		log.Fatal(err)
	}

	RunCmd(args[2:], envs)
}
