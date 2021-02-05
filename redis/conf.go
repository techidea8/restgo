package redis

import (
	"time"
)

type Conf struct {
	DB             int           `mapstructure:"db" json:"db" yaml:"db"`
	Addr           string        `mapstructure:"addr" json:"addr" yaml:"addr"`
	Password       string        `mapstructure:"password" json:"password" yaml:"password"`
	MaxIdle        int           `mapstructure:"maxIdle" json:"maxIdle" yaml:"max-idle"`
	IdleTimeoutSec time.Duration `mapstructure:"idleTimeoutSec" json:"idleTimeoutSec" yaml:"idle-timeout-sec"`
}
