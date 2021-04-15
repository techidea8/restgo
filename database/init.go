package database

import (
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

//@author: winlion
//@function: gormConfig
//@description: 根据配置决定是否开启日志
//@param: mod bool
//@return: *gorm.Config

func gormConfig(m Conf) *gorm.Config {
	//log.Debugf(" gormConfig %s", global.Config.Mysql.TablePrefix)

	var config = &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   m.TablePrefix, // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true,          // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	}
	//log.Println("%s", m.TablePrefix)
	return config
}

//@author: winlion
//@function: Gorm
//@description: 初始化数据库并产生数据库全局变量
//@return: *gorm.DB

func InitDataBase(m Conf, opts ...OrmOption) *gorm.DB {
	if m.DbType == DBTYPE_MYSQL {
		return InitMysql(m, opts...)
	} else if m.DbType == DBTYPE_SQLITE3 || m.DbType == DBTYPE_SQLITE {
		return InitSqlite3(m, opts...)
	} else if m.DbType == DBTYPE_POSTGRESQL {
		return InitPostGresSql(m, opts...)
	} else if m.DbType == DBTYPE_SQLSERVER {
		return InitSqlServer(m, opts...)
	} else {
		return InitSqlite3(m, opts...)
	}
}
