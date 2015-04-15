package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/l2x/wolffy/server/config"
)

var (
	DB orm.Ormer
)

func InitModels() error {
	//orm.Debug = true
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", "root:123456@/wolffy?charset=utf8")

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

	if config.NeedCreateAdministrator {
		err = UserModel.CheckCreateAdministor()
		if err != nil {
			return err
		}
	}

	return nil
}
