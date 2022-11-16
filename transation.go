package godb

import (
	"gorm.io/gorm"
)

type GormTransaction struct {
	QueryInterface
	instance *gorm.DB
}

func NewGormTransaction(instance *gorm.DB) TransactionInterface {
	return &GormTransaction{QueryInterface: NewGormQuery(instance), instance: instance}
}

func (r *GormTransaction) Commit() error {
	return r.instance.Commit().Error
}

func (r *GormTransaction) Rollback() error {
	return r.instance.Rollback().Error
}
