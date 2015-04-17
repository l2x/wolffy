package main

import (
	"flag"

	"github.com/l2x/wolffy/master/config"
	"github.com/l2x/wolffy/master/controllers"
	"github.com/l2x/wolffy/master/models"
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
