package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/l2x/wolffy/utils"
)

var (
	DatetimeFormat = "2006-01-02 15:04:05"
	DateFormat     = "2006-01-02"

	BasePath = "/tmp/"
	RepoPath = "/tmp/repo"
	DBPath   = "/tmp/data"

	PrivateKey = ""

	SessionInterval = 1
	SessionExpire   = 3600
	CookieName      = "wolffy_sessionid"
)

func InitConfig(conf string) error {
	err := loadPath()
	if err != nil {
		return err
	}

	err = loadConfig()
	if err != nil {
		return err
	}

	return nil
}

func loadConfig() error {
	PrivateKey = "123"

	return nil
}

func loadPath() error {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return err
	}

	BasePath = dir
	RepoPath = fmt.Sprintf("%s/%s", BasePath, "repo")
	DBPath = fmt.Sprintf("%s/%s", BasePath, "database")

	err = utils.Mkdir(BasePath, RepoPath, DBPath)
	if err != nil {
		return err
	}
	return nil
}
