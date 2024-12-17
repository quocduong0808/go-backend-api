package setting

type Config struct {
	Mysql  MysqlSetting  `mapstructure:"mysql"`
	Logger LoggerSetting `mapstructure:"logger"`
	Redis  RedisSetting  `mapstructure:"redis"`
}

type MysqlSetting struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	Dbname          string `mapstructure:"dbname"`
	MaxIdleCons     int    `mapstructure:"maxIdleCons"`
	MaxOpenCons     int    `mapstructure:"maxOpenCons"`
	ConnMaxLifeTime int    `mapstructure:"connMaxLifeTime"`
	AutoCreate      bool   `mapstructure:"autoCreate"`
}

type LoggerSetting struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"maxSize"`
	MaxBackups int    `mapstructure:"maxBackups"`
	MaxAge     int    `mapstructure:"maxAge"`
	Compress   bool   `mapstructure:"compress"`
}

type RedisSetting struct {
	Host           string `mapstructure:"host"`
	Port           int    `mapstructure:"port"`
	Password       string `mapstructure:"password"`
	Db             int    `mapstructure:"db"`
	PoolSize       int    `mapstructure:"poolSize"`
	MaxActiveConns int    `mapstructure:"maxActiveConns"`
	MaxIdleConns   int    `mapstructure:"maxIdleConns"`
}
