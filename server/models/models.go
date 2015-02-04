package models

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego/orm"
	"github.com/l2x/wolffy/server/config"
	_ "github.com/mattn/go-sqlite3"
)

var (
	DB orm.Ormer
)

func InitModels() error {
	dbPath := fmt.Sprintf("%s/%s", strings.TrimRight(config.DBPath, "/"), "data.db")

	orm.Debug = true
	orm.RegisterDriver("sqlite3", orm.DR_Sqlite)
	orm.RegisterDataBase("default", "sqlite3", dbPath)

	orm.RegisterModel(ClusterModel)
	orm.RegisterModel(DeployModel)
	orm.RegisterModel(DeployHistoryModel)
	orm.RegisterModel(MachineModel)
	orm.RegisterModel(ProjectModel)
	orm.RegisterModel(ProjectClusterModel)

	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		return err
	}

	DB = orm.NewOrm()

	return nil
}
