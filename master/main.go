package main

import (
	"flag"

	"github.com/l2x/wolffy/server/config"
	"github.com/l2x/wolffy/server/controllers"
	"github.com/l2x/wolffy/server/models"
)

var (
	configFile = ""
)

func init() {
	flag.StringVar(&configFile, "c", "config/config.ini", "config file")
	flag.Parse()

	err := config.InitConfig(configFile)
	if err != nil {
		panic(err)
	}
	err = controllers.InitControllers()
	if err != nil {
		panic(err)
	}
	err = models.InitModels()
	if err != nil {
		panic(err)
	}

}

func main() {

	router()
}
