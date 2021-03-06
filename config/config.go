package config

type Server struct {
	Addr       string `yaml:"addr"`
	Activation bool   `yaml:"activation"`  // activation strategy, true: automatically activate, false: activate by another request
	TokenExpire int   `yaml:"token-expire"`
	VCodeExpire int   `yaml:"vcode-expire"`
	LinkExpire int    `yaml:"link-expire"`
}

type Jwt struct {
	SecretKey  string `yaml:"secret-key"`
}

type MySQL struct {
	User       string `yaml:"user"`
	Password   string `yaml:"password"`
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
	DBName     string `yaml:"dbname"`
}

type Redis struct {
	Addr       string `yaml:"addr"`
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

type Config struct {
	Server     Server `yaml:"server"`
	Jwt        Jwt    `yaml:"jwt"`
	MySQL      MySQL  `yaml:"mysql"`
	Redis      Redis  `yaml:"redis"`
	Email      Email  `yaml:"email"`
	SMS        SMS    `yaml:"sms"`
}