package config

type Server struct {
	Port       int    `yaml:"port"`
	ApiVersion string `yaml:"api-version"`
	ApiPrefix  string `yaml:"api-prefix"`
}

type Jwt struct {
	SecretKey  string `yaml:"secret-key"`
}

type MySQL struct {
	UserName   string `yaml:"username"`
	Password   string `yaml:"password"`
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
	DBName     string `yaml:"dbname"`
}

type Redis struct {
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
}

type Email struct {
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
	Address    string `yaml:"address"`
	Password   string `yaml:"password"`
}

type SMS struct {
	API        string `yaml:"api"`
	Token      string `yaml:"token"`
}

type SystemConfig struct {
	Server     Server `yaml:"server"`
	Jwt        Jwt    `yaml:"jwt"`
	MySQL      MySQL  `yaml:"mysql"`
	Redis      Redis  `yaml:"redis"`
	Email      Email  `yaml:"email"`
	SMS        SMS    `yaml:"sms"`
}

type Config interface {
	ReadConfig(path string) (*SystemConfig, error)
	WriteConfig(path string, sys SystemConfig) error
}
