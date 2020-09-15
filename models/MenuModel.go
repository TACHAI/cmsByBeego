package models

import (
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"github.com/bitly/go-simplejson"
	_ "github.com/go-sql-driver/mysql"
	"sort"
)

type MenuModel struct {
	Mid int `orm:"pk;auto"`
	Parent int
	//Name string`orm:"size(45)"`
	Name string`orm:"size(45)"`
	Seq int
	//Format string `orm:"size(2048)";default({})`
	Format string `orm:"size(2048)";default({})`
}



type MenuTree struct {
	MenuModel
	Child[]MenuModel
}

func (m *MenuModel)TableName()string  {
	return "xcms_menu"
}

func MenuStruct()map[int]MenuTree  {
	query:=orm.NewOrm().QueryTable("xcms_menu")
	data:=make([]*MenuModel,0)
	query.OrderBy("parent","-seq").All(&data)

	var menu=make(map[int]MenuTree)
	if len(data)>0{
		for _,v:=range data{
			if 0==v.Parent{
				var tree = new(MenuTree)
				tree.MenuModel=*v
				menu[v.Mid]=*tree
			}else {
				if tmp,ok:=menu[v.Parent];ok{
					tmp.Child=append(tmp.Child,*v)
					menu[v.Parent]=tmp
				}
			}
		}
	}

	return menu
}

func MenuTreeStuct(user UserModel)map[int]MenuTree  {
	query:=orm.NewOrm().QueryTable("xcms_menu")
	data:=make([]*MenuModel,0)
	query.OrderBy("parent","-seq").Limit(1000).All(&data)

	var menu=make(map[int]MenuTree)
	//auth
	if len(user.AuthStr)>0{
		var authArr []int
		json.Unmarshal([]byte(user.AuthStr),&authArr)
		sort.Ints(authArr)

		for _,v:=range data{
			if 0==v.Parent{
				idx:=sort.Search(len(authArr), func(i int) bool {
					return authArr[i]>v.Mid
				})
				found:=(idx<len(authArr)&&authArr[idx]==v.Mid)
				if found{
					var tree=new(MenuTree)
					tree.MenuModel=*v
					menu[v.Mid]=*tree
				}
			}else{
				if tmp,ok:=menu[v.Parent];ok{
					tmp.Child=append(tmp.Child,*v)
					menu[v.Parent]=tmp
				}
			}

		}
	}
	return menu
}


func MenuList()([]*MenuModel,int64)  {
	query:=orm.NewOrm().QueryTable("xcms_menu")
	total,_:=query.Count()
	data:=make([]*MenuModel,0)
	query.OrderBy("parent","-seq").All(&data)
	return data,total
}

func ParentMenuList()[]*MenuModel  {
	query:=orm.NewOrm().QueryTable("xcms_menu").Filter("parent",0)
	data:=make([]*MenuModel,0)
	query.OrderBy("-seq").All(&data)

	return data
}

func MenuFormatStruct(mid int)*simplejson.Json  {
	menu:=MenuModel{Mid:mid}
	err:=orm.NewOrm().Read(&menu)
	if nil==err{
		jsonstruct,err2:=simplejson.NewJson([]byte(menu.Format))
		if(nil==err2){
			return jsonstruct
		}
	}
	return nil
}