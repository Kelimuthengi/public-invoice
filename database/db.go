package database

import (
	"fmt"

	"github.com/keliMuthengi/invoiving-api/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func init() {
	// env.LoadEnvVariable()
	env := env.NewEnv()
	host := env.Host
	user := env.User
	password := env.Password
	dbname := env.Dbname
	port := env.Port
	sslmode := env.Sslmode
	fmt.Println("PORT: < === > ", port)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s ", host, user, password, dbname, port, sslmode)
	// connectt to theos.Getenv("POSTGRES_URI") database:
	fmt.Println("dsn --->", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB = db

	// perform migrations
	for _, entity := range Entities {
		err := db.AutoMigrate(entity)
		if err != nil {
			panic(err)
		}
	}
	if err := db.Exec("ALTER TABLE users ALTER COLUMN email TYPE VARCHAR(255)").Error; err != nil {
		panic(err)
	}
}
