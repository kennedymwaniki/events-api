package pg

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func RunMigrations() {

	err := godotenv.Load()
    if err != nil {
       log.Fatal("Error loading .env file")
    }

	DATABASE_URL := os.Getenv("DATABASE_URL")
	m, err := migrate.New(
		"file://internal/infrastructure/pg/migrations/", DATABASE_URL)
	if err != nil {
		fmt.Printf("Error creating migration instance: %v\n", err)
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		// log.Fatal(err)
		fmt.Printf("Error running migrations😭: %v\n", err)
		return
	}
	fmt.Println("Migrations ran successfully🥳")
}