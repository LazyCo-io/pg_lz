package actions

import (
	"log"

	"github.com/LazyCo-io/pg_lz/pkg/postgres"
	"github.com/spf13/cobra"
)

func Upgrade(cmd *cobra.Command, args []string) {
	server := postgres.NewServer("postgresql://postgres:postgres@127.0.0.1:5432/postgres")

	err := server.Upgrade()
	if err != nil {
		log.Fatal(err)
	}

}
