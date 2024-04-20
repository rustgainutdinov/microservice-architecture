package main

import (
	"io"
	"log"
	"os"
	"usercrud/src/infrastructure"
)

func main() {
	rootDir := "/app/data/migrations/"
	migrations, err := os.ReadDir(rootDir)
	if err != nil {
		log.Fatal(err)
	}

	sqlHandler := infrastructure.NewSqlHandler()

	var file *os.File
	for _, m := range migrations {
		file, err = os.Open(rootDir + m.Name())
		if err != nil {
			log.Fatal(err)
		}

		b, _ := io.ReadAll(file)
		sqlHandler.Execute(string(b))
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	infrastructure.Init(sqlHandler)
}
