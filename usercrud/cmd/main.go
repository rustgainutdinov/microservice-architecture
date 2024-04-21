package main

import (
	"log"
	"usercrud/src/domain"
	"usercrud/src/infrastructure"
)

func main() {
	sqlHandler := infrastructure.NewSqlHandler()

	err := sqlHandler.Migrate(domain.User{})
	if err != nil {
		log.Fatal(err)
	}

	infrastructure.Init(sqlHandler)
}
