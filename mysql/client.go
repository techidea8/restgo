package mysql

import "github.com/techidea8/restgo"

//@author: winlion
//@function: Gorm
//@description: 初始化数据库并产生数据库全局变量
//@return: *gorm.DB

func InitDatabase(m Conf,opts ...OrmOption) *gorm.DB {
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Addr + ")/" + m.Dbname + "?" + m.Query
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         250,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置

	}
	cfg := gormConfig()
	for _,opt := range opts{
		opt(cfg)
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig),cfg); err != nil {

		log.Info("mysql start error:", err.Error())
		os.Exit(0)
		return nil
	} else {
		sqlDB, err := db.DB()

		if err != nil {
			log.Error("mysql start error", err.Error())
			os.Exit(0)
			return nil
		}

		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		log.Info("mysql start ok")

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
type OrmOption func(*gorm.Config)


//@author: winlion
//@function: gormConfig
//@description: 根据配置决定是否开启日志
//@param: mod bool
//@return: *gorm.Config

func gormConfig() *gorm.Config {
	log.Debugf(" gormConfig %s", global.Config.Mysql.TablePrefix)

	var config = &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   global.Config.Mysql.TablePrefix, // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true,                            // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	}
	log.Debugf("%s", global.Config.Mysql.TablePrefix)
	return config
}
