package scotch

import (
	"database/sql"
	"log"

	"github.com/CloudyKit/jet/v6"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/prateekjoshi2013/scotch/render"
)

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
	Routes   *chi.Mux
	Render   *render.Render
	JetViews *jet.Set
	Session  *scs.SessionManager
	DB       Database
	config   config // private scotch framework config
}

type config struct {
	port        string
	renderer    string
	cookie      cookieConfig
	sessionType string
	database    databaseConfig
}

type cookieConfig struct {
	name     string
	lifetime string
	persist  string // whether the cookie should be persisted across requests (between browser sessions)
	secure   string // whether the cookie should be sent with encryoption
	domain   string
}

type databaseConfig struct {
	dsn      string
	database string
}

type Database struct {
	DatabaseType string
	Pool         *sql.DB
}
