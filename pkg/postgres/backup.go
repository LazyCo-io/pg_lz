package postgres

import (
	"fmt"
	"log"
	"time"
)

const (
	backupLabelPrefix = "pg_lz_backup_label"
)

func (s *Server) startBackup() error {
	var backupLSN string

	backupLabel := fmt.Sprintf("%s_%s", backupLabelPrefix, time.Now().Format("20060102T150405"))
	log.Println("creating backup with label:", backupLabel)
	err := s.conn.QueryRow(s.ctx, `SELECT pg_catalog.pg_backup_start($1)`, backupLabel).Scan(&backupLSN)
	if err != nil {
		return fmt.Errorf("could not run pg_start_backups: %w", err)
	}
	log.Println("backup created! lsn:", backupLSN)

	return nil
}

func (s *Server) stopBackup() error {
	log.Println("stopping the backup")
	_, err := s.conn.Exec(s.ctx, `SELECT pg_catalog.pg_backup_stop()`)
	if err != nil {
		return fmt.Errorf("could not run pg_start_backups: %w", err)
	}
	log.Println("backup stoped!")

	return nil
}
