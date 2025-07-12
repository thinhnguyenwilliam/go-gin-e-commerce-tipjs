package setting

type Config struct {
	Server ServerSetting `mapstructure:"server"`
	MySQL  MySQLSetting  `mapstructure:"mysql"`
	Logger LoggerSetting `mapstructure:"log" yaml:"log"`
	Redis  RedisSetting  `mapstructure:"redis" yaml:"redis"`
}

type RedisSetting struct {
	Addr     string `mapstructure:"addr" yaml:"addr"`
	Password string `mapstructure:"password" yaml:"password"`
	DB       int    `mapstructure:"db" yaml:"db"`
	PoolSize int    `mapstructure:"poolsize" yaml:"poolsize"`
}

type MySQLSetting struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	User            string `mapstructure:"user"`
	Password        string `mapstructure:"password"`
	DBName          string `mapstructure:"dbname"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int    `mapstructure:"connexMaxLifetime"`
}

type LoggerSetting struct {
	LogLevel    string `mapstructure:"logLevel" yaml:"logLevel"`
	FileLogName string `mapstructure:"fileLogName" yaml:"fileLogName"`
	MaxSize     int    `mapstructure:"maxSize" yaml:"maxSize"`
	MaxBackups  int    `mapstructure:"maxBackups" yaml:"maxBackups"`
	MaxAge      int    `mapstructure:"maxAge" yaml:"maxAge"`
	Compress    bool   `mapstructure:"compress" yaml:"compress"`
}

type ServerSetting struct {
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}
