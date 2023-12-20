package scotch

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/prateekjoshi2013/scotch/render"
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

	// create loggers
	infoLog, errorLog := s.startLoggers()
	s.InfoLog = infoLog
	s.ErrorLog = errorLog
	s.Debug = s.Debug || os.Getenv("DEBUG") == "true"

	s.config = config{
		port:     os.Getenv("PORT"),
		renderer: os.Getenv("RENDERER"),
	}

	// set version
	s.Version = VERSION

	// set root path
	s.RootPath = rootPath

	// create routes
	s.Routes = s.routes().(*chi.Mux)

	// create render engine
	s.CreateRenderer()
	
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

func (s *Scotch) startLoggers() (*log.Logger, *log.Logger) {
	var infoLog *log.Logger
	var errorLog *log.Logger

	infoLog = log.New(os.Stdout, "INFO: \t", log.Ldate|log.Ltime|log.Lshortfile)
	errorLog = log.New(os.Stderr, "ERROR: \t", log.Ldate|log.Ltime|log.Lshortfile)
	return infoLog, errorLog
}

func (s *Scotch) ListenAndServe() {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", s.config.port),
		ErrorLog:     s.ErrorLog,
		Handler:      s.Routes,
		IdleTimeout:  30 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 600 * time.Second,
	}

	s.InfoLog.Printf("Listening on port %s", s.config.port)
	err := srv.ListenAndServe()
	s.ErrorLog.Fatal(err)
}

func (s *Scotch) CreateRenderer() {
	renderer := render.Render{
		Renderer: s.config.renderer,
		RootPath: s.RootPath,
		Port:     s.config.port,
	}
	s.Render = &renderer
}
