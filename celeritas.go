package celeritas

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const version = "1.0.0"

type Celeritas struct {
	AppName  string
	Debug    bool
	Version  string
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	RootPath string
	// config won't be exported because the users shouldn't have access to the configs
	config   Config
}

// config of the package celeritas
type Config struct {
	port string 
	renderer string
}


func (c *Celeritas) New(rootPath string) error {
	pathConfig := initPaths{
		rootPath:    rootPath,
		folderNames: []string{"handlers", "migrations", "views", "data", "public", "tmp", "logs", "middleware"},
	}

	err := c.Init(pathConfig)
	if err != nil {
		return err
	}

	// verify existence of .env file
	err = c.checkDotEnv(rootPath)
	if err != nil {
		return err
	}

	// if present then read .env
	err = godotenv.Load(rootPath + "/.env")
	if err != nil {
		return err
	}

	// logger
	infoLog, errorLog := c.startLoggers()
	c.InfoLog = infoLog
	c.ErrorLog = errorLog
	c.Version = version
	// everything that comes from env file is a string
	// the string is being converted to bool ignoring the error
	c.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	c.RootPath = rootPath

	// configs
	c.config = Config{
		port: os.Getenv("PORT"),
		renderer: os.Getenv("RENDERER"),
	}
	return nil
}

func (c *Celeritas) Init(p initPaths) error {
	root := p.rootPath
	for _, path := range p.folderNames {
		// create folder if it doesn't exist
		err := c.CreateDirIfNotExist(root + "/" + path)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Celeritas) checkDotEnv(path string) error {
	err := c.CreateFileIfNotExists(fmt.Sprintf("%s/.env", path))
	if err != nil {
		return err
	}
	return nil
}

func (c *Celeritas) startLoggers() (*log.Logger, *log.Logger) {
	var infoLog *log.Logger
	var errorLog *log.Logger

	infoLog = log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stdout, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile)

	return infoLog, errorLog
}
