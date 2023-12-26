package scotch

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)


func (s *Scotch) MigrateUp(dsn string) error {
	m, err := migrate.New("file://"+s.RootPath+"/migrations", dsn)
	if err != nil {
		return err
	}
	defer m.Close()
	if err := m.Up(); err != nil {
		log.Println("error running migrations: ", err)
		return err
	}
	return nil
}

func (s *Scotch) MigrateDownAll(dsn string) error {
	m, err := migrate.New("file://"+s.RootPath+"/migrations", dsn)
	if err != nil {
		return err
	}
	defer m.Close()
	if err := m.Down(); err != nil {
		log.Println("error running migrations: ", err)
		return err
	}
	return nil
}

func (s *Scotch) Steps(n int, dsn string) error {
	m, err := migrate.New("file://"+s.RootPath+"/migrations", dsn)
	if err != nil {
		return err
	}
	defer m.Close()
	if err := m.Steps(n); err != nil {
		return err
	}
	return nil
}

func (s *Scotch) MigrateForce(dsn string) error {
	m, err := migrate.New("file://"+s.RootPath+"/migrations", dsn)
	if err != nil {
		return err
	}
	defer m.Close()
	if err := m.Force(-1); err != nil {
		return err
	}
	return nil
}
