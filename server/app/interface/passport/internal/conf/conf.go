package conf

var config *Config

type Config struct {
	Server  ServerConfig  `yaml:"server" json:"server"`
	Log     LogConfig     `yaml:"log" json:"log"`
	Time    TimeConfig    `yaml:"time" json:"time"`
	Service ServiceConfig `yaml:"service" json:"service"`
	Jwt     JwtConfig     `yaml:"jwt" json:"jwt"`
	Github  GithubConfig  `yaml:"github" json:"github"`
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

type ServiceItem struct {
	Host string `yaml:"host" json:"host"`
	Port uint64 `yaml:"port" json:"port"`
}

type ServiceConfig struct {
	Administrator ServiceItem `yaml:"administrator" json:"administrator"`
	Article       ServiceItem `yaml:"article" json:"article"`
}

type JwtItemData struct {
	Key string `yaml:"key"`
	Iv  uint64 `yaml:"iv"`
}

type JwtConfig struct {
	Administrator  JwtItemData `yaml:"administrator"`
	Visitor        JwtItemData `yaml:"visitor"`
	VisitorSession JwtItemData `yaml:"visitor_session"`
}

type GithubConfig struct {
	ClientID     string `yaml:"client_id" json:"client_id"`
	ClientSecret string `yaml:"client_secret" json:"client_secret"`
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

func GetServiceConfig() *ServiceConfig {
	return &config.Service
}

func GetJwtConfig() *JwtConfig {
	return &config.Jwt
}

func GetGithubConfig() *GithubConfig {
	return &config.Github
}
