package handlers

import (
	"fmt"

	"github.com/lpernett/godotenv"
)

func LoadEnvVariable() {

	err := godotenv.Load()

	if err != nil {
		fmt.Println("env_loading_error", err.Error())
		panic("Error loading env files:")
	}
}
