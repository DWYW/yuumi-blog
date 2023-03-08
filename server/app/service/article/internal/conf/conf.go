package conf

var config *Config

type Config struct {
	Server     ServerConfig     `yaml:"server" json:"server"`
	Log        LogConfig        `yaml:"log" json:"log"`
	Time       TimeConfig       `yaml:"time" json:"time"`
	Pagination PaginationConfig `yaml:"pagination" json:"pagination"`
	Mysql      MysqlConfig      `yaml:"mysql" json:"mysql"`
}

type ServerConfig struct {
	Http HTTPServerConfig `yaml:"http" json:"http"`
	Grpc GRPCServerConfig `yaml:"grpc" json:"grpc"`
}

type HTTPServerConfig struct {
	Host string               `yaml:"host" json:"host"`
	Port uint64               `yaml:"port" json:"port"`
	Cors HTTPServerCorsConfig `yaml:"cors" json:"cors"`
}

type HTTPServerCorsConfig struct {
	Enabled         bool   `yaml:"enabled" json:"enabled"`
	WhiteListRegexp string `yaml:"white_list_regexp" json:"white_list_regexp"`
}

type GRPCServerConfig struct {
	Host string `yaml:"host" json:"host"`
	Port uint64 `yaml:"port" json:"port"`
}

type LogConfig struct {
	Dir          string `yaml:"dir"`
	RotationHour uint64 `yaml:"rotation_hour"`
	MaxAgeDay    uint64 `yaml:"max_age_day"`
}

type TimeConfig struct {
	Location string `yaml:"location" json:"location"`
}

type PaginationConfig struct {
	PageSize int64 `yaml:"page_size" json:"page_size"`
}

type MysqlConfig struct {
	Host     string          `yaml:"host" json:"host"`
	Port     uint64          `yaml:"port" json:"port"`
	Username string          `yaml:"username" json:"username"`
	Password string          `yaml:"password" json:"password"`
	Dbname   string          `yaml:"dbname" json:"dbname"`
	Charset  string          `yaml:"charset" json:"charset"`
	Pool     MysqlPoolConfig `yaml:"pool" json:"pool"`
}

type MysqlPoolConfig struct {
	MaxIdleConns uint64 `yaml:"max_idle_conns" json:"max_idle_conns"`
	MaxOpenConns uint64 `yaml:"max_open_conns" json:"max_open_conns"`
	MaxLifetime  uint64 `yaml:"max_lifetime" json:"max_lifetime"`
}

func Init(value *Config) {
	config = value
}

func GetConfig() *Config {
	return config
}

func GetHTTPServerConfig() *HTTPServerConfig {
	return &config.Server.Http
}

func GetHTTPCorsConfig() *HTTPServerCorsConfig {
	return &config.Server.Http.Cors
}

func GetGRPCServerConfig() *GRPCServerConfig {
	return &config.Server.Grpc
}

func GetLogConfig() *LogConfig {
	return &config.Log
}

func GetTimeConfig() *TimeConfig {
	return &config.Time
}

func GetPaginationConfig() *PaginationConfig {
	return &config.Pagination
}

func GetMysqlConfig() *MysqlConfig {
	return &config.Mysql
}
