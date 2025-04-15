package bootstrap

import (
	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

type Env struct {
	ServerHost     string `mapstructure:"SERVER_HOST"`
	ServerPort     string `mapstructure:"SERVER_PORT"`
	ContextTimeout int    `mapstructure:"CONTEXT_TIMEOUT"`
	DBHost         string `mapstructure:"DB_HOST"`
	DBPort         string `mapstructure:"DB_PORT"`
	DBUser         string `mapstructure:"DB_USER"`
	DBPass         string `mapstructure:"DB_PASS"`
	DBName         string `mapstructure:"DB_NAME"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Can't find the file .env : %v", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatalf("Environment can't be loaded: %v", err)
	}

	log.Info("Environment variables loaded successfully")
	return &env
}
