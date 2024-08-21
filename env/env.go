package env

import (
	"fmt"
	"log"

	"github.com/lpernett/godotenv"
	"github.com/spf13/viper"
)

type Env struct {
	Port         string `mapstructure:"PORT"`
	PostgressURI string `mapstructure:"POSTGRES_URI"`
	JwtKey       string `mapstructure:"JWT_KEY"`
	Host         string `mapstructure:"HOST"`
	User         string `mapstructure:"USER"`
	Password     string `mapstructure:"PASSWORD"`
	Dbname       string `mapstructure:"DB_NAME"`
	Sslmode      string `mapstructure:"SSL_MODE"`
	Mode         string `mapstructure:"MODE"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	neterr := viper.Unmarshal(&env)

	if neterr != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	return &env
}

func LoadEnvVariable() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("env_loading_error", err.Error())
		panic("Error loading env files:")
	}
}
