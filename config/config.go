package config


type MySqlConfig struct {
	User string
	Password string
	Host string
	DBName string
}

type HttpConfig struct {
	Addr string
}