package cmd

import (
	"encoding/json"
	"os"

	"github.com/spf13/cobra"
)

var (
	release   = "UNKNOWN"
	buildDate = "UNKNOWN"
	gitHash   = "UNKNOWN"
)

func version() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "app version",
		RunE: func(cmd *cobra.Command, args []string) error {
			return json.NewEncoder(os.Stdout).Encode(struct {
				Release   string
				BuildDate string
				GitHash   string
			}{
				Release:   release,
				BuildDate: buildDate,
				GitHash:   gitHash,
			})
		},
	}
}
