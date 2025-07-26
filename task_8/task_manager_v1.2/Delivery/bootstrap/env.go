package bootstrap

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv         string `mapstructure:"APP_ENV"`
	ServerAddress  string `mapstructure:"SERVER_ADDRESS"`
	MongoUri       string `mapstructure:"MONGO_URI"`
	DBName         string `mapstructure:"DB_NAME"`
	RefTS          string `mapstructure:"REFRESH_TOKEN_SECRET"`
	AccTS          string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefTE          int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccTE          int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	ContextTimeout int    `mapstructure:"CTX_TIMEOUT"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}

	return &env
}
