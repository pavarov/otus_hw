package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	timeout := flag.Duration("timeout", time.Second*10, "timeout in seconds for connection to server")
	flag.Parse()

	args := flag.Args()
	if len(args) < 2 || len(args[0]) == 0 || len(args[1]) == 0 {
		log.Fatal("invalid arguments")
	}

	address := net.JoinHostPort(args[0], args[1])
	client := NewTelnetClient(address, *timeout, os.Stdin, os.Stdout)
	if err := client.Connect(); err != nil {
		log.Fatal("failed to create connection: ", err)
	}
	defer func(client TelnetClient) {
		err := client.Close()
		if err != nil {
			log.Fatal("failed to close client")
		}
	}(client)

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	go func() {
		if err := client.Send(); err != nil {
			log.Fatal("failed to send message by the client")
		}
	}()

	go func() {
		if err := client.Receive(); err != nil {
			log.Fatal("failed to receive message by the client")
		}
	}()

	<-ctx.Done()
}
