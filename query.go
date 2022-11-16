package godb

import (
	"errors"
	"gorm.io/gorm"
)

type GormQuery struct {
	instance *gorm.DB
}

func NewGormQuery(instance *gorm.DB) QueryInterface {
	return &GormQuery{instance}
}

func (r *GormQuery) Driver() Driver {
	return Driver(r.instance.Dialector.Name())
}

func (r *GormQuery) Count(count *int64) error {
	return r.instance.Count(count).Error
}

func (r *GormQuery) Create(value interface{}) error {
	return r.instance.Create(value).Error
}

func (r *GormQuery) Delete(value interface{}, conds ...interface{}) error {
	return r.instance.Delete(value, conds...).Error
}

func (r *GormQuery) Distinct(args ...interface{}) QueryInterface {
	tx := r.instance.Distinct(args...)

	return NewGormQuery(tx)
}

func (r *GormQuery) Exec(sql string, values ...interface{}) error {
	return r.instance.Exec(sql, values...).Error
}

func (r *GormQuery) Find(dest interface{}, conds ...interface{}) error {
	return r.instance.Find(dest, conds...).Error
}

func (r *GormQuery) First(dest interface{}) error {
	err := r.instance.First(dest).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}

	return err
}

func (r *GormQuery) FirstOrCreate(dest interface{}, conds ...interface{}) error {
	var err error
	if len(conds) > 1 {
		err = r.instance.Attrs([]interface{}{conds[1]}...).FirstOrCreate(dest, []interface{}{conds[0]}...).Error
	} else {
		err = r.instance.FirstOrCreate(dest, conds...).Error
	}

	return err
}

func (r *GormQuery) ForceDelete(value interface{}, conds ...interface{}) error {
	return r.instance.Unscoped().Delete(value, conds...).Error
}

func (r *GormQuery) Get(dest interface{}) error {
	return r.instance.Find(dest).Error
}

func (r *GormQuery) Group(name string) QueryInterface {
	tx := r.instance.Group(name)

	return NewGormQuery(tx)
}

func (r *GormQuery) Having(query interface{}, args ...interface{}) QueryInterface {
	tx := r.instance.Having(query, args...)

	return NewGormQuery(tx)
}

func (r *GormQuery) Join(query string, args ...interface{}) QueryInterface {
	tx := r.instance.Joins(query, args...)

	return NewGormQuery(tx)
}

func (r *GormQuery) Limit(limit int) QueryInterface {
	tx := r.instance.Limit(limit)

	return NewGormQuery(tx)
}

func (r *GormQuery) Model(value interface{}) QueryInterface {
	tx := r.instance.Model(value)

	return NewGormQuery(tx)
}

func (r *GormQuery) Offset(offset int) QueryInterface {
	tx := r.instance.Offset(offset)

	return NewGormQuery(tx)
}

func (r *GormQuery) Order(value interface{}) QueryInterface {
	tx := r.instance.Order(value)

	return NewGormQuery(tx)
}

func (r *GormQuery) OrWhere(query interface{}, args ...interface{}) QueryInterface {
	tx := r.instance.Or(query, args...)

	return NewGormQuery(tx)
}

func (r *GormQuery) Pluck(column string, dest interface{}) error {
	return r.instance.Pluck(column, dest).Error
}

func (r *GormQuery) Raw(sql string, values ...interface{}) QueryInterface {
	tx := r.instance.Raw(sql, values...)

	return NewGormQuery(tx)
}

func (r *GormQuery) Save(value interface{}) error {
	return r.instance.Save(value).Error
}

func (r *GormQuery) Scan(dest interface{}) error {
	return r.instance.Scan(dest).Error
}

func (r *GormQuery) Select(query interface{}, args ...interface{}) QueryInterface {
	tx := r.instance.Select(query, args...)

	return NewGormQuery(tx)
}

func (r *GormQuery) Table(name string, args ...interface{}) QueryInterface {
	tx := r.instance.Table(name, args...)

	return NewGormQuery(tx)
}

func (r *GormQuery) Update(column string, value interface{}) error {
	return r.instance.Update(column, value).Error
}

func (r *GormQuery) Updates(values interface{}) error {
	return r.instance.Updates(values).Error
}

func (r *GormQuery) Where(query interface{}, args ...interface{}) QueryInterface {
	tx := r.instance.Where(query, args...)

	return NewGormQuery(tx)
}

func (r *GormQuery) WithTrashed() QueryInterface {
	tx := r.instance.Unscoped()

	return NewGormQuery(tx)
}

func (r *GormQuery) Scopes(funcs ...func(QueryInterface) QueryInterface) QueryInterface {
	var gormFuncs []func(*gorm.DB) *gorm.DB
	for _, item := range funcs {
		gormFuncs = append(gormFuncs, func(db *gorm.DB) *gorm.DB {
			item(&GormQuery{db})

			return db
		})
	}

	tx := r.instance.Scopes(gormFuncs...)

	return NewGormQuery(tx)
}

