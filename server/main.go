package main

import (
	"github.com/l2x/wolffy/server/config"
	"github.com/l2x/wolffy/server/controllers"
	"github.com/l2x/wolffy/server/models"
)

func init() {
	err := config.InitConfig("")
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
