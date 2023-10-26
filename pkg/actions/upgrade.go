package actions

import (
	"fmt"
	"log"

	"github.com/LazyCo-io/pg_lz/pkg/postgres"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Upgrade(cmd *cobra.Command, args []string) {
	err := postgres.NewServer(buildConnString()).Upgrade()
	if err != nil {
		log.Fatal(err)
	}

}

func buildConnString() string {
	return fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s",
		viper.GetString("username"),
		viper.GetString("password"),
		viper.GetString("host"),
		viper.GetString("port"),
		viper.GetString("dbname"),
	)
}
