package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

type conf struct {
	DBDriver       string `mapstructure:"DB_DRIVER"`
	DBHost         string `mapstructure:"DB_HOST"`
	DBPort         string `mapstructure:"DB_PORT"`
	DBUser         string `mapstructure:"DB_USER"`
	DBPassword     string `mapstructure:"DB_PASSWORD"`
	DBName         string `mapstructure:"DB_NAME"`
	WebServicePort string `mapstructure:"WEB_SERVICE_PORT"`
	JwtSecret      string `mapstructure:"JWT_SECRET"`
	JwtExpiresIn   int    `mapstructure:"JWT_EXPIRESIN"`
	TokenAuth      jwtauth.JWTAuth
}

func LoadConfig(path string) (*conf, error) {
	var cfg conf

	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&cfg)

	if err != nil {
		panic(err)
	}

	cfg.TokenAuth = *jwtauth.New("HS256", []byte(cfg.JwtSecret), nil)

	return &cfg, nil
}
