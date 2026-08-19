package postgresql

import (
	"context"
	"fmt"
	"learning-project/internal/config"
	"learning-project/pkg/logging"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewClient(ctx context.Context, postgreConfig *config.PostgreSQLConfig) (db *pgxpool.Pool, err error) {

	logger := logging.GetLogger()

	connectionStr := postgreConfig.DSN()

	poolCfg, err := pgxpool.ParseConfig(connectionStr)

	if err != nil {
		return nil, fmt.Errorf("Failed Parse Postgre Config: %w", err)
	}

	poolCfg.MaxConns = int32(postgreConfig.MaxOpenConns)
	poolCfg.MinConns = int32(postgreConfig.MaxIdleConns)
	poolCfg.MaxConnLifetime = postgreConfig.ConnMaxLifetime
	poolCfg.MaxConnIdleTime = postgreConfig.ConnMaxIdleTime

	poolCfg.ConnConfig.ConnectTimeout = postgreConfig.ConnectTimeout

	logger.Info("Connection Postgre DB")

	pool, err := pgxpool.NewWithConfig(ctx, poolCfg)

	if err != nil {
		return nil, fmt.Errorf("Failed to connect postgresql: %w", err)
	}

	if err = pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("Failed ping to Postgre DB: %w", err)
	}

	return pool, nil
}
