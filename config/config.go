package config

type Config struct {
	Host           string `mapstructure:"DB_HOST"`
	Port           string `mapstructure:"DB_PORT"`
	Debug          string `mapstructure:"APP_DEBUG"`
	Username       string `mapstructure:"DB_USERNAME"`
	Password       string `mapstructure:"DB_PASSWORD"`
	Name           string `mapstructure:"DB_NAME"`
	MaxIdle        int    `mapstructure:"DB_MAX_IDLE"`
	MaxConn        int    `mapstructure:"DB_MAX_CONN"`
	MaxIdleTimeout string `mapstructure:"DB_MAX_IDLE_TIMEOUT"`
	AppPort        string `mapstructure:"APP_PORT"`
	JwtSecret      string `mapstructure:"JWT_SECRET"`
}
