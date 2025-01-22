package config

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Cluster struct {
	Conn *pgx.Conn
}

func NewCluster(ctx context.Context) (*Cluster, error) {
	dsn := fmt.Sprintf(
		"postgres://localhost:5432?dbname=warehouse&user=postgres&password=example&sslmode=disable",
	)

	fmt.Printf("dsn: %s\n", dsn)
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("pgx.Connect: %w", err)
	}

	return &Cluster{Conn: conn}, nil
}
