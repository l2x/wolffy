package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	orm.RegisterDriver("sqlite3", orm.DR_Sqlite)
	orm.RegisterDataBase("default", "sqlite3", "database/wolffy.db")

	orm.RegisterModel(new(Project))
}
