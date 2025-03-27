package postgres

type Config struct {
	Host     string `env:"HOST" yaml:"HOST"`
	Port     string `env:"PORT" yaml:"PORT"`
	Name     string `env:"NAME" yaml:"NAME"`
	User     string `env:"USER" yaml:"USER"`
	Password string `env:"PASSWORD" yaml:"PASSWORD"`
}
