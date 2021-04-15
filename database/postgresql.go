package database

import (
	"gorm.io/gorm"
)

//@author: winlion
//@function: Gorm
//@description: 初始化数据库并产生数据库全局变量
//@return: *gorm.DB
func InitPostGresSql(m Conf, opts ...OrmOption) *gorm.DB {
	return nil
}
