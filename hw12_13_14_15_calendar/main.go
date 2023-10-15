package main

import (
	"context"

	_ "github.com/lib/pq"
	"github.com/pavarov/otus_hw/hw12_13_14_15_calendar/cmd"
)

func main() {
	ctx := context.Background()
	cmd.Execute(ctx)
}
