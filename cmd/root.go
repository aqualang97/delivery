package cmd

import (
	"database/sql"
	config "delivery/configs"
	"github.com/aqualang97/logger/v4"
	"github.com/spf13/cobra"
)

func CreateRootCmd(cfg *config.Config, conn *sql.DB, myLogger *logger.Logger) *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "demo",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	}
	rootCmd.AddCommand(CreateCreateCmd(cfg, conn, myLogger))
	rootCmd.AddCommand(CreateServerCmd(cfg, conn, myLogger))
	rootCmd.AddCommand(CreateTruncateCmd(cfg, conn))
	return rootCmd
}
