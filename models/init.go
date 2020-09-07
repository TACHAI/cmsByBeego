package models

import "github.com/astaxie/beego/orm"

func init()  {
	orm.RegisterDataBase("default","mysql","rooot:hzc778209(127.0.0.1:3306)/xcms?charset=utf-8")
	orm.RegisterModel(new(MenuModel))
}
