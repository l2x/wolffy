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

func init() {
	dbPath := fmt.Sprintf("%s/%s", strings.TrimRight(config.DBPath, "/"), "data.db")

	orm.RegisterDriver("sqlite3", orm.DR_Sqlite)
	orm.RegisterDataBase("default", "sqlite3", dbPath)

	orm.RegisterModel(ProjectModel)
	orm.RegisterModel(ClusterModel)
	orm.RegisterModel(DeployModel)
	orm.RegisterModel(ProductModel)

	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		panic(err)
	}

	DB = orm.NewOrm()
}
