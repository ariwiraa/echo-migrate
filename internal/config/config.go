package config

import (
	"log"

	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

//config holds configuration fo the projet
type Config struct {
	Port           string `env:"PORT,default=6666"`
	Env            string `env:"ENV,default=development"`
	Database       DatabaseConfig
	JWTConfig      JWTConfig
	InternalConfig InternalConfig
}

//DatabaseConfig holds configuration for database
type DatabaseConfig struct {
	Host     string `env:"DATABASE_HOST,default=localhost"`
	Port     string `env:"DATABASE_PORT,default=5432"`
	Username string `env:"DATABASE_USERNAME,required"`
	Password string `env:"DATABASE_PASSWORD,required"`
	Name     string `env:"DATABASE_NAME,required"`
}

type JWTConfig struct {
	SecretKey string `env:"JWT_SECRET_KEY,required"`
}

type InternalConfig struct {
	Username string `env:"SVC_USERNAME,required"`
	Password string `env:"SVC_PASSWORD,required"`
}

func NewConfig(env string) (*Config, error) {
	var config Config
	if err := godotenv.Load(env); err != nil {
		log.Println(errors.Wrap(err, "[NewConfig] error reading .env file, defaulting to OS environment variables"))
	}

	if err := envdecode.Decode(&config); err != nil {
		return nil, errors.Wrap(err, "[NewConfig] error decoding me")
	}

	return &config, nil
}
