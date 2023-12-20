package scotch

type initPaths struct {
	rootPath    string
	folderNames []string
}

type Scotch struct {
	AppName string
	Debug   bool
	Version string
}