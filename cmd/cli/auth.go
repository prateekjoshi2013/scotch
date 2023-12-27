package main

import (
	"fmt"
	"log"
	"time"
)

func doAuth() error {
	// migrations
	dbType := sco.DB.DatabaseType
	fileName := fmt.Sprintf("%d_create_auth_tables", time.Now().UnixMicro())
	upFile := sco.RootPath + "/migrations/" + fileName + ".up.sql"
	downFile := sco.RootPath + "/migrations/" + fileName + ".down.sql"
	log.Println(dbType, upFile, downFile)
	err := copyFileFromTemplate("templates/migrations/auth_tables."+dbType+".sql", upFile)
	if err != nil {
		exitGraceFully(err)
	}
	err = copyDataToFile([]byte(`
	drop table if exists users cascade;
	drop table if exists tokens cascade;
	drop table if exists remember_tokens cascade;
	`), downFile)
	if err != nil {
		exitGraceFully(err)
	}
	// run migrations
	doMigrate("up", "")
	if err != nil {
		exitGraceFully(err)
	}
	// copy files over
	err = copyFileFromTemplate("templates/data/user.go.txt", sco.RootPath+"/data/user.go")
	if err != nil {
		exitGraceFully(err)
	}
	err = copyFileFromTemplate("templates/data/token.go.txt", sco.RootPath+"/data/token.go")
	if err != nil {
		exitGraceFully(err)
	}
	return nil
}
