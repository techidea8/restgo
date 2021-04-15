package database

type Conf struct {
	DbType       string `mapstructure:"db-type" json:"dbType" yaml:"db-type"`
	Addr         string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Dbname       string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`
	Username     string `mapstructure:"username" json:"username" yaml:"username"`
	Query        string `mapstructure:"query" json:"query" yaml:"query"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"`
	TablePrefix  string `mapstructure:"table-prefix" json:"tablePrefix" yaml:"table-prefix"`
	Debug        bool   `mapstructure:"debug" json:"debug" yaml:"debug"`
}
