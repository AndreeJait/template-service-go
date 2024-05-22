package config

type Config struct {
	Service Service `json:"service" yaml:"service"`

	DB DB `json:"db" yaml:"db"`
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
