package config

type Config struct {
	Service Service `json:"service" yaml:"service"`

	DB DB `json:"db" yaml:"db"`

	Secret Secret `json:"secret" yaml:"secret"`

	Setting Setting `json:"setting" yaml:"setting"`
}

type Setting struct {
	JwtTokenDurationExpiredInMinute        int `json:"jwt_token_duration_expired_in_minute" yaml:"jwt_token_duration_expired_in_minute"`
	JwtRefreshTokenDurationExpiredInMinute int `json:"jwt_refresh_token_duration_expired_in_minute" yaml:"jwt_refresh_token_duration_expired_in_minute"`
}

type Secret struct {
	Jwt  string `json:"jwt" yaml:"jwt"`
	Hash string `json:"hash" yaml:"hash"`
}

type DB struct {
	Host     string `json:"host" yaml:"host"`
	Port     string `json:"port" yaml:"port"`
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
	DBName   string `json:"db_name" yaml:"db_name"`
}

type Service struct {
	Name string `json:"name" yaml:"name"`
	Port string `json:"port" yaml:"port"`
	Host string `json:"host" yaml:"host"`
	Url  string `json:"url" yaml:"url"`
}
