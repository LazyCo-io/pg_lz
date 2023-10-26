package postgres

import (
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func (s *Server) Upgrade() error {
	log.Println("connecting to the database...")

	var err error
	s.conn, err = pgx.Connect(s.ctx, s.ConnString)
	if err != nil {
		return fmt.Errorf("could not connect in the database: %w", err)
	}
	defer func() {
		s.conn.Close(s.ctx)
		log.Println("disconnecting...")
	}()

	err = s.startBackup()
	if err != nil {
		return fmt.Errorf("could not start the backup: %w", err)
	}

	// copia x N

	err = s.stopBackup()
	if err != nil {
		return fmt.Errorf("could not stop the backup: %w", err)
	}

	return nil
}
