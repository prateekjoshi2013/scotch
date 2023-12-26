package main

import (
	"errors"
	"fmt"
	"time"
)

func doMake(arg2, arg3 string) error {
	switch arg2 {
	case "migration":
		dbType := sco.DB.DatabaseType
		if arg3 == "" {
			exitGraceFully(errors.New("you must give migration a name"))
		}
		fileName := fmt.Sprintf("%d_%s", time.Now().UnixMicro(), arg3)
		upFile := sco.RootPath + "/migrations/" + fileName + "." + dbType + ".up.sql"
		downFile := sco.RootPath + "/migrations/" + fileName + "." + dbType + ".down.sql"
		err := copyFileFromTemplate("templates/migrations/migration."+dbType+".up.sql", upFile)
		if err != nil {
			exitGraceFully(err)
		}
		err = copyFileFromTemplate("templates/migrations/migration."+dbType+".down.sql", downFile)
		if err != nil {
			exitGraceFully(err)
		}
	}
	return nil
}
