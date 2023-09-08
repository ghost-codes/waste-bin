package config

import (
	"fmt"
	"time"
)

type Config struct {
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	FirebaseConfigPath   string        `mapstructure:"FIREBASE_CONFIG_PATH"`
	SecretKey            string        `mapstructure:"SECRET_KEY"`
	DBDriver             string        `mapstructure:"DB_DRIVER"`
	ServerAddr           string        `mapstructure:"SEVER_ADDR"`
	DBHOST               string        `mapstructure:"DB_HOST"`
	DBPort               string        `mapstructure:"DB_PORT"`
	DBUser               string        `mapstructure:"DB_USER"`
	DBPassword           string        `mapstructure:"DB_PASSWORD"`
	DBName               string        `mapstructure:"DB_NAME"`
	GRPCServerAddress    string        `mapstructure:"GRPC_SEVER_ADDR"`
	HttpServerAddress    string        `mapstructure:"HTTP_SEVER_ADDR"`
	ProjectAPIKey        string        `mapstructure:"PROJECT_API_KEY"`
	AppEnv               string        `mapstructure:"APP_ENV"`
	DevUID               string        `mapstructure:"DEV_UID"`
}

func (config *Config) DBSource() string {
	return fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=disable", config.DBUser, config.DBPassword, config.DBHOST, config.DBPort, config.DBName)
}
