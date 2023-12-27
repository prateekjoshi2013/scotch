package main

import (
	"fmt"
	"time"
)

func doSessionTable() error {

	dbType := sco.DB.DatabaseType
	if dbType == "mariadb" {
		dbType = "mysql"
	}

	if dbType == "postgresql" {
		dbType = "postgres"
	}

	fileName := fmt.Sprintf("%d_create_sessions_table", time.Now().UnixMicro())

	upFile := sco.RootPath + "/migrations/" + fileName + "." + dbType + ".up.sql"
	downFile := sco.RootPath + "/migrations/" + fileName + "." + dbType + ".down.sql"
	err := copyFileFromTemplate("templates/migrations/"+dbType+"_session.sql", upFile)
	if err != nil {
		return err
	}
	err = copyDataToFile([]byte("drop table sessions;"), downFile)
	if err != nil {
		exitGraceFully(err)
	}
	err = doMigrate("up", "")
	if err != nil {
		exitGraceFully(err)
	}
	return nil
}
