package storage

import (
	"context"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/fx"
)

var Module = fx.Module("storage",
	fx.Provide(New),
)

type Storage struct {
	db *sqlx.DB
}

func New(ctx context.Context) (*Storage, error) {
	db, err := sqlx.ConnectContext(ctx, "sqlite3", "./db/local.sqlite")
	if err != nil {
		return nil, err
	}
	return &Storage{
		db: db,
	}, nil
}

func (s *Storage) Execute(ctx context.Context, query string, args ...interface{}) error {
	_, err := s.db.ExecContext(ctx, query, args)
	return err
}

func (s *Storage) Select(ctx context.Context, query string, resp interface{}, args ...interface{}) error {
	return s.db.SelectContext(ctx, &resp, query, args...)
}

func (s *Storage) Find(ctx context.Context, query string, resp interface{}, args ...interface{}) error {
	return s.db.GetContext(ctx, &resp, query, args...)
}
