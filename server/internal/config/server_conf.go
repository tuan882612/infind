package config

type server struct {
	Address string `env:"SERVER_ADDR,required"`
}
