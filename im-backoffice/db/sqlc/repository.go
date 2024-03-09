package db

import (
	"database/sql"
)

type Repository interface {
	Querier
}

type SQLRepository struct {
	*Queries
	db *sql.DB
}

func NewRepository(conn *sql.DB) Repository {
	return &SQLRepository{
		db:      conn,
		Queries: New(conn),
	}
}

// transaction related
//func (r *SQLRepository) execTx(ctx context.Context, fn func(*Queries) error) error {
//	tx, err := r.db.BeginTx(ctx, nil)
//	if err != nil {
//		return err
//	}
//	q := New(tx)
//	err = fn(q)
//	if err != nil {
//		if rbErr := tx.Rollback(); rbErr != nil {
//			return fmt.Errorf("tx err: %v, rb: err: %v", err, rbErr)
//		}
//		return err
//	}
//	return tx.Commit()
//}
