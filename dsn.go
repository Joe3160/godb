package godb

import (
	"fmt"
	"github.com/goravel/framework/facades"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func getSqlserverGormConfig(connection string) gorm.Dialector {
	host := facades.Config.GetString("database.connections." + connection + ".host")
	if host == "" {
		return nil
	}
	port := facades.Config.GetString("database.connections." + connection + ".port")
	database := facades.Config.GetString("database.connections." + connection + ".database")
	username := facades.Config.GetString("database.connections." + connection + ".username")
	password := facades.Config.GetString("database.connections." + connection + ".password")
	dsn:= fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s&encrypt=disable",
		username, password, host, port, database)
	return sqlserver.New(sqlserver.Config{
		DSN: dsn,
	})
}

func getMysqlGormConfig(connection string) gorm.Dialector {
	host := facades.Config.GetString("database.connections." + connection + ".host")
	if host == "" {
		return nil
	}
	port := facades.Config.GetString("database.connections." + connection + ".port")
	database := facades.Config.GetString("database.connections." + connection + ".database")
	username := facades.Config.GetString("database.connections." + connection + ".username")
	password := facades.Config.GetString("database.connections." + connection + ".password")
	charset := facades.Config.GetString("database.connections." + connection + ".charset")
	loc := facades.Config.GetString("database.connections." + connection + ".loc")
	dsn:= fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s",
		username, password, host, port, database, charset, true, loc)
	return mysql.New(mysql.Config{
		DSN: dsn,
	})
}
func getPostgresqlGormConfig(connection string) gorm.Dialector {
	host := facades.Config.GetString("database.connections." + connection + ".host")
	if host == "" {
		return nil
	}
	port := facades.Config.GetString("database.connections." + connection + ".port")
	database := facades.Config.GetString("database.connections." + connection + ".database")
	username := facades.Config.GetString("database.connections." + connection + ".username")
	password := facades.Config.GetString("database.connections." + connection + ".password")
	sslmode := facades.Config.GetString("database.connections." + connection + ".sslmode")
	timezone := facades.Config.GetString("database.connections." + connection + ".timezone")
	dsn:= fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host, username, password, database, port, sslmode, timezone)
	return postgres.New(postgres.Config{
		DSN: dsn,
	})
}

func getSqliteGormConfig(connection string) gorm.Dialector {
	dsn := facades.Config.GetString("database.connections." + connection + ".database")
	if dsn == "" {
		return nil
	}
	return sqlite.Open(dsn)
}
