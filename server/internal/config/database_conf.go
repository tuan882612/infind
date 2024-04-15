package config

type database struct {
	PgURL string `env:"PG_URL,required"`
}
