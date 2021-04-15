package database

import (
	"gorm.io/gorm"
)

var (
	DbEngin *gorm.DB
)

const (
	DBTYPE_SQLITE3    = "sqlite3"
	DBTYPE_SQLITE    = "sqlite"
	DBTYPE_MYSQL      = "mysql"
	DBTYPE_POSTGRESQL = "postgresql"
	DBTYPE_SQLSERVER  = "sqlserver"
)

type OrmOption func(*gorm.Config)
