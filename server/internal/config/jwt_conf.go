package config

type jwt struct {
	Secret string `env:"JWT_SECRET,required"`
}
