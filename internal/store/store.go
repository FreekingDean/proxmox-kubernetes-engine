package store

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/FreekingDean/proxmox-kubernetes-engine/internal/logger"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/fx"
)

type DB interface {
	ExecContext(context.Context, string, ...any) (sql.Result, error)
	SelectContext(context.Context, any, string, ...any) error
	GetContext(context.Context, any, string, ...any) error
}

type Store struct {
	db     DB
	logger logger.Logger
}

type StoreParams struct {
	fx.In

	Logger logger.Logger
}

func New(ctx context.Context, p StoreParams) (*Store, error) {
	db, err := sqlx.ConnectContext(ctx, "sqlite3", "file:./db/local.sqlite?_foreign_keys=on")
	if err != nil {
		return nil, err
	}
	_, err = db.Exec("PRAGMA foriegn_keys = ON")
	if err != nil {
		return nil, err
	}
	return &Store{
		db:     db,
		logger: p.Logger,
	}, nil
}

func (s *Store) Execute(ctx context.Context, query string, args ...interface{}) error {
	_, err := s.db.ExecContext(ctx, query, args...)
	return err
}

func (s *Store) Select(ctx context.Context, query string, resp any, args ...interface{}) error {
	return s.db.SelectContext(ctx, resp, query, args...)
}

func (s *Store) Find(ctx context.Context, query string, resp interface{}, args ...interface{}) error {
	return s.db.GetContext(ctx, resp, query, args...)
}

type txfunc func(*Store) error
type TxStarter interface {
	BeginTxx(context.Context, *sql.TxOptions) (*sqlx.Tx, error)
}

func (s *Store) Transaction(ctx context.Context, f txfunc) error {
	txdb, ok := s.db.(TxStarter)
	if !ok {
		return fmt.Errorf("cannot start tx in context")
	}

	tx, err := txdb.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Commit()
	}()

	txs := &Store{
		db:     tx,
		logger: s.logger,
	}
	if err = f(txs); err != nil {
		return err
	}
	return tx.Commit()
}
