package godb

import (
	"errors"
	"fmt"
	"github.com/goravel/framework/facades"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"log"
	"os"
	"runtime"
	"time"
)

const (
	Mysql      = "mysql"
	Postgresql = "postgresql"
	Sqlite     = "sqlite"
	Sqlserver  = "sqlserver"
	MsSQL      = "mssql"
)

func getGormConfig(connection string) (gorm.Dialector, error) {
	driver := facades.Config.GetString("database.connections." + connection + ".driver")

	switch driver {
	case Mysql:
		return getMysqlGormConfig(connection), nil
	case Postgresql:
		return getPostgresqlGormConfig(connection), nil
	case Sqlite:
		return getSqliteGormConfig(connection), nil
	case Sqlserver, MsSQL:
		return getSqlserverGormConfig(connection), nil
	default:
		return nil, errors.New(fmt.Sprintf("err database driver: %s, only support mysql, postgresql, sqlite and sqlserver", driver))
	}
}

func NewGormInstance(connection string) (*gorm.DB, error) {
	gormConfig, err := getGormConfig(connection)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("init gorm config error: %v", err))
	}
	if gormConfig == nil {
		return nil, nil
	}

	var logLevel gormLogger.LogLevel
	if facades.Config.GetBool("app.debug") {
		logLevel = gormLogger.Info
	} else {
		logLevel = gormLogger.Error
	}
	var colourful bool
	if runtime.GOOS == "linux" {
		colourful = true
	}
	logger := New(log.New(os.Stdout, "\r\n", log.LstdFlags), gormLogger.Config{
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  logLevel,
		IgnoreRecordNotFoundError: true,
		Colorful:                  colourful,
	})

	return gorm.Open(gormConfig, &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		SkipDefaultTransaction:                   true,
		Logger:                                   logger.LogMode(logLevel),
	})
}
