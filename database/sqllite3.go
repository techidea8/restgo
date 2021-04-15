package database

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//@author: winlion
//@function: Gorm
//@description: 初始化数据库并产生数据库全局变量
//@return: *gorm.DB

func InitSqlite3(m Conf, opts ...OrmOption) *gorm.DB {
	cfg := gormConfig(m)
	for _, opt := range opts {
		opt(cfg)
	}

	if db, err := gorm.Open(sqlite.Open(m.Addr+string(os.PathSeparator)+m.Dbname), cfg); err != nil {
		log.Fatal("sqlite3 start error:", err.Error(), m.Addr+string(os.PathSeparator)+m.Dbname)
		os.Exit(0)
		return nil
	} else {
		DbEngin = db
		return db
	}
}
