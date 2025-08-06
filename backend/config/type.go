package config

type Config struct {
	AppName  string         `yaml:"appName"`
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"db"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

type DatabaseConfig struct {
	Driver     string `yaml:"driver"`
	ConnString string `yaml:"connString"`
}
