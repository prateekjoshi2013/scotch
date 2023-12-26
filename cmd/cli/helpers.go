package main

import (
	"os"

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
