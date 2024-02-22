package database

type Config struct {
	User     string `mapstructure:"user" env:"USER"`
	Password string `mapstructure:"password" env:"PASSWORD"`
	Host     string `mapstructure:"host" env:"HOST"`
	Name     string `mapstructure:"name" env:"NAME"`
}
