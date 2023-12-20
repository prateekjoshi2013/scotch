package scotch

import (
	"fmt"

	"github.com/joho/godotenv"
)

const VERSION = "1.0.0"

func (s *Scotch) New(rootPath string) error {
	pathConfig := initPaths{
		rootPath: rootPath,
		folderNames: []string{
			"handlers",
			"migrations",
			"views",
			"data",
			"public",
			"tmp",
			"logs",
			"middlewares",
		},
	}

	err := s.Init(pathConfig)
	if err != nil {
		return err
	}

	err = s.checkDotEnv(rootPath)
	if err != nil {
		return err
	}

	// read .env

	err = godotenv.Load(rootPath + "/.env")
	if err != nil {
		return err
	}

	return nil
}

func (s *Scotch) Init(p initPaths) error {
	root := p.rootPath
	for _, path := range p.folderNames {
		// create folder if it does not exist
		err := s.CreateDirIfNotExist(root + "/" + path)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Scotch) checkDotEnv(path string) error {
	err := s.CreateFileIfNotExist(fmt.Sprintf("%s/.env", path))
	if err != nil {
		return err
	}
	return nil
}
