package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/iancoleman/strcase"
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
	case "auth":
		err := doAuth()
		if err != nil {
			exitGraceFully(err)
		}
	case "handler":
		if arg3 == "" {
			exitGraceFully(errors.New("you must give a handler a name"))
		}
		fileName := sco.RootPath + "/handlers/" + strings.ToLower(arg3) + ".go"
		if fileExists(fileName) {
			exitGraceFully(errors.New("handler already exists"))
		}
		data, err := templateFS.ReadFile("templates/handlers/handler.go.txt")
		if err != nil {
			exitGraceFully(err)
		}
		handler := string(data)
		handler = strings.ReplaceAll(handler, "$HANDLERNAME$", strcase.ToCamel(arg3))
		err = os.WriteFile(fileName, []byte(handler), 0644)
		if err != nil {
			exitGraceFully(err)
		}

	}
	return nil
}
