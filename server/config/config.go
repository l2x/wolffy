package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Unknwon/goconfig"
	"github.com/l2x/wolffy/utils"
)

var (
	DatetimeFormat = "2006-01-02 15:04:05"
	DateFormat     = "2006-01-02"

	NeedCreateAdministrator = false

	config     *goconfig.ConfigFile
	ConfigFile = "config/config.ini"
	BasePath   = "/tmp/"
	RepoPath   = ""
	DBPath     = ""

	Port = ":9020"

	PrivateKey = ""

	SessionInterval = 1
	SessionExpire   = 3600
	CookieName      = "wolffy_sid"
)

func InitConfig() error {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return err
	}
	cf := fmt.Sprintf("%s/%s", dir, ConfigFile)

	err = loadConfig(cf)
	if err != nil {
		return err
	}

	err = getParams()
	if err != nil {
		return err
	}

	err = goconfig.SaveConfigFile(config, cf)
	if err != nil {
		return err
	}

	return nil
}

func loadConfig(cf string) error {
	f, err := os.OpenFile(cf, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	f.Close()

	config, err = goconfig.LoadConfigFile(cf)
	if err != nil {
		return err
	}
	return nil
}

func getParams() error {
	// path
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return err
	}

	if BasePath, err = config.GetValue("", "basePath"); err != nil || BasePath == "" {
		BasePath = dir
		config.SetValue("", "basePath", BasePath)
	}
	if RepoPath, err = config.GetValue("", "repoPath"); err != nil || RepoPath == "" {
		RepoPath = fmt.Sprintf("%s/%s", BasePath, "repo")
		config.SetValue("", "repoPath", RepoPath)
	}

	if DBPath, err = config.GetValue("", "dbPath"); err != nil || DBPath == "" {
		DBPath = fmt.Sprintf("%s/%s", BasePath, "database")
		config.SetValue("", "dbPath", DBPath)
	}

	err = utils.Mkdir(BasePath, RepoPath, DBPath)
	if err != nil {
		return err
	}

	// privatekey
	if PrivateKey, err = config.GetValue("", "privateKey"); err != nil || PrivateKey == "" {
		uuid, err := utils.UUID()
		if err != nil {
			return err
		}
		PrivateKey = uuid
		config.SetValue("", "privateKey", PrivateKey)
		NeedCreateAdministrator = true
	}

	// port
	if port, err := config.GetValue("", "port"); err == nil || port == "" {
		Port = port
	}
	config.SetValue("", "port", Port)

	return nil
}
