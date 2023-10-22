package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/exp/slog"
)

var (
	configFile string

	rootCmd = &cobra.Command{
		Use:   "calendar",
		Short: "calendar app",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Usage()
		},
	}
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&configFile, "config", ".env", "config file (default is $HOME/.env)")
}

func initConfig() {
	viper.SetConfigFile(configFile)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("config file is not registered: %s", err))
	}
}

func Execute(ctx context.Context) {
	rootCmd.AddCommand(version())
	rootCmd.AddCommand(server(ctx))
	rootCmd.AddCommand(schedulerCommand(ctx))
	rootCmd.AddCommand(senderCommand(ctx))
	if err := rootCmd.Execute(); err != nil {
		slog.Error("failed to execute root command", "error:", err)
		os.Exit(1)
	}
}
