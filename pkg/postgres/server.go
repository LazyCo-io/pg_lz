package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Server struct {
	ConnString string
	ctx        context.Context
	conn       *pgx.Conn
}

func NewServer(connStr string) *Server {
	return &Server{
		ConnString: connStr,
		ctx:        context.Background(),
	}
}
