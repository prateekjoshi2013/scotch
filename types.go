package scotch

import "log"

type initPaths struct {
	rootPath    string
	folderNames []string
}

type Scotch struct {
	AppName  string
	Debug    bool
	Version  string
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	RootPath string
	config config // private scotch framework config
}

type config struct {
	port     string
	renderer string
}
