package godb

import (
	"context"
	"gorm.io/gorm"
)

type GormDB struct {
	QueryInterface
	instance *gorm.DB
}

func (r *GormDB) Begin() (TransactionInterface, error) {
	tx := r.instance.Begin()
	return NewGormTransaction(tx), tx.Error
}

func NewGormDB(ctx context.Context, connection string) (DBInterface, error) {
	db, err := NewGormInstance(connection)
	if err != nil {
		return nil, err
	}
	if db == nil {
		return nil, nil
	}

	if ctx != nil {
		db = db.WithContext(ctx)
	}

	return &GormDB{
		QueryInterface: NewGormQuery(db),
		instance:       db,
	}, nil
}

