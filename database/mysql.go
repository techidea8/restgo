package database

import (
	"os"

	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//@author: winlion
//@function: Gorm
//@description: 初始化数据库并产生数据库全局变量
//@return: *gorm.DB

func InitMysql(m Conf, opts ...OrmOption) *gorm.DB {
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Addr + ")/" + m.Dbname + "?" + m.Query
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         250,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置

	}
	cfg := gormConfig(m)
	for _, opt := range opts {
		opt(cfg)
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), cfg); err != nil {

		log.Fatal("mysql start error:", err.Error())
		os.Exit(0)
		return nil
	} else {
		sqlDB, err := db.DB()

		if err != nil {
			log.Fatal("mysql start error", err.Error())
			os.Exit(0)
			return nil
		}

		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		log.Println("mysql start ok")

		if m.Debug {
			db.Logger = logger.Default.LogMode(logger.Info)
			db = db.Debug()
		} else {
			db.Logger = logger.Default.LogMode(logger.Error)
		}
		DbEngin = db
		return db
	}
}
