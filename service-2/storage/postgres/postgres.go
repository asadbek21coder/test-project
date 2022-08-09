package postgres

import (
	"context"

	"github.com/asadbek21coder/test-project/service-2/config"
	"github.com/asadbek21coder/test-project/service-2/storage"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	db        *pgxpool.Pool
	service_2 storage.Service_2_I
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

func (s *Store) Service_2_I() storage.Service_2_I {
	if s.service_2 == nil {
		s.service_2 = NewService_2_Repo(s.db)
	}

	return s.service_2
}
