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

	config   *goconfig.ConfigFile
	BasePath = "/tmp/"
	RepoPath = ""

	Port = ":9020"

	PrivateKey = ""

	SessionInterval = 1
	SessionExpire   = 3600
	CookieName      = "wolffy_sid"

	DBHost = ""
	DBUser = ""
	DBPwd  = ""
	DBName = ""
)

func InitConfig(configFile string) error {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return err
	}
	cf := fmt.Sprintf("%s/%s", dir, configFile)

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
	var err error

	// db
	if DBHost, err = config.GetValue("database", "dbHost"); err != nil {
		return err
	}
	if DBName, err = config.GetValue("database", "dbName"); err != nil {
		return err
	}
	if DBUser, err = config.GetValue("database", "dbUser"); err != nil {
		return err
	}
	if DBPwd, err = config.GetValue("database", "dbPwd"); err != nil {
		return err
	}

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

	err = utils.Mkdir(BasePath, RepoPath)
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
	}

	// port
	if port, err := config.GetValue("", "port"); err != nil && port != "" {
		Port = port
	}
	config.SetValue("", "port", Port)

	return nil
}
