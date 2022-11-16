package godb

import (
	"context"
)

type OrmInterface interface {
	Connection(name string) OrmInterface
	Query() DBInterface
	Transaction(txFunc func(tx TransactionInterface) error) error
	WithContext(ctx context.Context) OrmInterface
}

//go:generate mockery --name=DB
type DBInterface interface {
	QueryInterface
	Begin() (TransactionInterface, error)
}

//go:generate mockery --name=Transaction
type TransactionInterface interface {
	QueryInterface
	Commit() error
	Rollback() error
}

type QueryInterface interface {
	Driver() Driver
	Count(count *int64) error
	Create(value interface{}) error
	Delete(value interface{}, conds ...interface{}) error
	Distinct(args ...interface{}) QueryInterface
	Exec(sql string, values ...interface{}) error
	Find(dest interface{}, conds ...interface{}) error
	First(dest interface{}) error
	FirstOrCreate(dest interface{}, conds ...interface{}) error
	ForceDelete(value interface{}, conds ...interface{}) error
	Get(dest interface{}) error
	Group(name string) QueryInterface
	Having(query interface{}, args ...interface{}) QueryInterface
	Join(query string, args ...interface{}) QueryInterface
	Limit(limit int) QueryInterface
	Model(value interface{}) QueryInterface
	Offset(offset int) QueryInterface
	Order(value interface{}) QueryInterface
	OrWhere(query interface{}, args ...interface{}) QueryInterface
	Pluck(column string, dest interface{}) error
	Raw(sql string, values ...interface{}) QueryInterface
	Save(value interface{}) error
	Scan(dest interface{}) error
	Scopes(funcs ...func(QueryInterface) QueryInterface) QueryInterface
	Select(query interface{}, args ...interface{}) QueryInterface
	Table(name string, args ...interface{}) QueryInterface
	Update(column string, value interface{}) error
	Updates(values interface{}) error
	Where(query interface{}, args ...interface{}) QueryInterface
	WithTrashed() QueryInterface
}


