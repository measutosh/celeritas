package celeritas

const version = "1.0.0"

type Celeritas struct {
	AppName string
	Debug   bool
	Version string
}


func (c *Celeritas) New(rootPath string) error {
	pathConfig := initPaths{
		rootPath: rootPath,
		folderNames: []string{"handlers", "migrations", "views", "data", "logs", "tmp", "middleware"},
	}

	err:= c.Init(pathConfig)
	if err != nil {
		return nil
	}

	return nil
}

func (c *Celeritas) Init(p initPaths) error {
	root := p.rootPath
	for _, path := range p.folderNames {
		// create the folder if doesn't exit
		err := c.CreateDirIfNotExist(root + "/" + path)
		if err != nil {
			return err
		}
	}	
	return nil
}