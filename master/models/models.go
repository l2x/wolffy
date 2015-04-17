package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/l2x/wolffy/master/config"
)

var (
	DB orm.Ormer
)

func InitModels() error {
	//orm.Debug = true
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", config.DBUser, config.DBPwd, config.DBHost, config.DBName))

	orm.RegisterModel(ClusterModel)
	orm.RegisterModel(ClusterNodeModel)
	orm.RegisterModel(DeployModel)
	orm.RegisterModel(DeployHistoryModel)
	orm.RegisterModel(NodeModel)
	orm.RegisterModel(ProjectModel)
	orm.RegisterModel(ProjectClusterModel)
	orm.RegisterModel(UserModel)

	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		return err
	}

	DB = orm.NewOrm()

	err = UserModel.CheckCreateAdministor()
	if err != nil {
		return err
	}

	return nil
}
