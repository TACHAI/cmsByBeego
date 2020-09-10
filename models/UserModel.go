package models

import "github.com/astaxie/beego/orm"

type UserModel struct {
	UserId int `orm:"pk;auto"`
	Userkey string `"size(45)"`
	Name string`orm:"size(45)"`
	Password string `orm:"size(255)"`
	AuthStr string `orm:"size(255)"`
}


func (m *UserModel)TableName()string  {
	return "xcms_user"
}


func UserList(pageSize,page int)([]*UserModel,int64)  {
	query:=orm.NewOrm().QueryTable("xcms_user")
	total,_:=query.Count()
	data:=make([]*UserModel,0)
	query.OrderBy("-uid","-seq").All(&data)
	return data,total
}

func GetUserByName(userkey string)UserModel  {
	user:=UserModel{Userkey:userkey}
	o:=orm.NewOrm()
	o.Read(&user)
	return user
}