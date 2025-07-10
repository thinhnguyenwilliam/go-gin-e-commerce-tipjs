package setting

type Config struct {
	Server ServerSetting `mapstructure:"server"`
	MySQL  MySQLSetting  `mapstructure:"mysql"`
	Logger LoggerSetting `mapstructure:"log" yaml:"log"`
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
