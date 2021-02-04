package restgo

import (
	"net/http"
	"github.com/techidea8/restgo/redis"
	"github.com/techidea8/restgo/mysql"
)
var (
	DbEngin *gorm.DB
	RedisEngin *RedisClient
)