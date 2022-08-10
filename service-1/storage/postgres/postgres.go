package postgres

import (
	"context"

	"github.com/asadbek21coder/test-project/service-1/config"
	"github.com/asadbek21coder/test-project/service-1/storage"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	db        *pgxpool.Pool
	service_1 storage.Service_1_I
}

func NewPostgres(psqlConnString string, cfg config.Config) (storage.StorageI, error) {
	config, err := pgxpool.ParseConfig(psqlConnString)
	if err != nil {
		return nil, err
	}

	config.AfterConnect = nil
	config.MaxConns = int32(cfg.PostgresMaxConnections)

	pool, err := pgxpool.ConnectConfig(context.Background(), config)

	return &Store{
		db: pool,
	}, err
}

func (s *Store) Service_1_I() storage.Service_1_I {
	if s.service_1 == nil {
		s.service_1 = NewService_1_Repo(s.db)
	}

	return s.service_1
}
