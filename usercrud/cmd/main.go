package main

import (
	"fmt"
	"log"
	"usercrud/src/domain"
	"usercrud/src/infrastructure"

	"github.com/caarlos0/env/v11"
)

type config struct {
	DBHost     string `env:"DB_HOST"`
	DBPort     int    `env:"DB_PORT" envDefault:"3306"`
	DBName     string `env:"DB_NAME"`
	DBUserName string `env:"DB_USER_NAME"`
	DBPass     string `env:"DB_PASS"`
}

func main() {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}
	sqlHandler := infrastructure.NewSqlHandler(cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBUserName, cfg.DBPass)

	err := sqlHandler.Migrate(domain.User{})
	if err != nil {
		log.Fatal(err)
	}

	infrastructure.Init(sqlHandler)
}
