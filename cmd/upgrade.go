/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"os/user"

	"github.com/LazyCo-io/pg_lz/pkg/actions"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// upgradeCmd represents the upgrade command
var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrades a PostgreSQL instance with logical replication",
	// todo
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	Run: actions.Upgrade,
}

func init() {
	rootCmd.AddCommand(upgradeCmd)

	upgradeCmd.Flags().StringP("host", "h", getEnvOrDefault("PGHOST", "localhost"), "database server host or socket directory")
	upgradeCmd.Flags().StringP("port", "p", getEnvOrDefault("PGPORT", "5432"), "database server port")
	upgradeCmd.Flags().StringP("username", "U", getEnvOrDefault("PGUSER", getUser()), "database user name")
	upgradeCmd.Flags().StringP("password", "W", getEnvOrDefault("PGPASSWORD", ""), "password")
	upgradeCmd.Flags().StringP("dbname", "d", getEnvOrDefault("PGDATABASE", getUser()), "database name to connect to")

	viper.BindPFlag("host", upgradeCmd.Flags().Lookup("host"))
	viper.BindPFlag("port", upgradeCmd.Flags().Lookup("port"))
	viper.BindPFlag("username", upgradeCmd.Flags().Lookup("username"))
	viper.BindPFlag("password", upgradeCmd.Flags().Lookup("password"))
	viper.BindPFlag("dbname", upgradeCmd.Flags().Lookup("dbname"))
}

// TODO: descobrir o jeito de fazer com o viper
func getEnvOrDefault(varName, defaultVal string) string {
	out := os.Getenv(varName)
	if len(out) == 0 {
		return defaultVal
	}

	return out

}

func getUser() string {
	user, err := user.Current()
	if err != nil {
		return "postgres"
	}

	return user.Username
}
