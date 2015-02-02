package config

import (
	"fmt"
	"os"
)

var (
	DatetimeFormat = "2006-01-02 15:04:05"
	DateFormat     = "2006-01-02"

	BasePath string = "/tmp/repo"
	DBPath   string = "/tmp/data"
)

func InitConfig(conf string) error {
	fmt.Println(BasePath)
	err := os.MkdirAll(BasePath, 0755)
	if err != nil {
		return err
	}
	err = os.MkdirAll(DBPath, 0755)
	if err != nil {
		return err
	}

	return nil
}
