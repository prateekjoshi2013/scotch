package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"github.com/prateekjoshi2013/scotch"
)

func setup() {
	err := godotenv.Load()
	if err != nil {
		exitGraceFully(err)
	}

	path, err := os.Getwd()
	if err != nil {
		exitGraceFully(err)
	}
	sco = scotch.Scotch{RootPath: path}
	sco.DB.DatabaseType = os.Getenv("DATABASE_TYPE")
}

func getDSN() string {
	dbType := sco.DB.DatabaseType
	if dbType == "pgx" {
		dbType = "postgres"
	}
	if dbType == "postgres" {
		var dsn string
		if os.Getenv("DATABASE_PASS") != "" {
			dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
				os.Getenv("DATABASE_USER"),
				os.Getenv("DATABASE_PASS"),
				os.Getenv("DATABASE_HOST"),
				os.Getenv("DATABASE_PORT"),
				os.Getenv("DATABASE_NAME"),
				os.Getenv("DATABASE_SSL_MODE"),
			)
		} else {
			dsn = fmt.Sprintf("postgres://%s@%s:%s/%s?sslmode=%s",
				os.Getenv("DATABASE_USER"),
				os.Getenv("DATABASE_HOST"),
				os.Getenv("DATABASE_PORT"),
				os.Getenv("DATABASE_NAME"),
				os.Getenv("DATABASE_SSL_MODE"),
			)
		}
		return dsn
	} else {
		return "mysql://" + sco.BuildDSN()
	}

}

func showHelp() {
	color.Yellow(`
	Available commands:
	
	help	- Show this help
	version	- Show version
	migrate	- runs all up migrations that have not been run previously
	migrate down	- reverses the most recent migration
	migrate reset	- runs all down migrations in reverse order, and then all up migrations
	make migation <name>	- Creates new migration up and down with the given name
	make auth    - Creates models, migrations, tables for auth system
	make handler <name>    - Creates new handler with the given name in handlers directory
	make model <name>    - Creates new model with the given name in data directory
	`)
}
