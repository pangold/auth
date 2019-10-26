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
	Port       string `yaml:"port"`
	DBName     string `yaml:"dbname"`
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
	Email      Email  `yaml:"email"`
	SMS        SMS    `yaml:"sms"`
}

type Config interface {
	ReadConfig() (*SystemConfig, error)
	WriteConfig(sys SystemConfig) error
}
