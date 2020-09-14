package controllers

import (
	"cmsByBeego/models"
	"github.com/astaxie/beego/orm"
	"github.com/bitly/go-simplejson"
	"strconv"
	"time"
)

type DataController struct {
	BaseController
	Mid int
}

func (c *DataController)Prepare()  {
	c.BaseController.Prepare()

	midstr:=c.Ctx.Input.Param(":mid")
	mid,err:=strconv.Atoi(midstr)
	if nil==err &&mid>0{
		c.Mid=mid
	}else {
		c.setTpl()
	}

}

func (c *DataController)Index()  {
	sj:=models.MenuFormatStruct(c.Mid)
	if nil==sj{
		title:=make(map[string]string)
		titlemap:=sj.Get("schema")
		for k,_:=range titlemap.MustMap(){
			stype:=titlemap.GetPath(k,"type").MustString()
			if "object"!=stype&&"array"!=stype{
				if len(titlemap.GetPath(k,"title").MustString())>0{
				title[k] = titlemap.GetPath(k, "title").MustString()

				}else {
					title[k]=k
			}
			}
		}
		c.Data["Title"]=title
	}
	c.LayoutSections=make(map[string]string)
	c.LayoutSections["footerjs"]="data/footerjs.html"
	c.setTpl()
}

func (c *DataController)List()  {
	page,err:=c.GetInt("page")
	if err!=nil{
		page=1
	}
	size,err:=c.GetInt("limit")
	if err!=nil{
		size=20
	}
	data,total:=models.DataList(c.Mid,size,page)
	c.listJsonResult(200,"ok",total,data)
}

func (c *DataController)Edit()  {
	did,_:=c.GetInt("did")
	if did>0{
		c.Data["Did"]=did
	}

	c.initForm(did)

	c.LayoutSections=make(map[string]string)
	c.LayoutSections["footerjs"]="data/footerjs_edit.html"
	c.setTpl("data/edit.hrml","common/layout_edit.html")
}

func (c *DataController)initForm(did int)  {
	 format:=models.MenuFormatStruct(c.Mid)
	 if nil ==format{
		 return
	 }
	 schemaMap:=format.Get("schema")
	 formArray:=format.Get("form")

	 one:=models.DataRead(did)
	 if nil !=one{
	 	for k,_:=range schemaMap.MustMap(){
			switch schemaMap.GetPath(k,"type").MustString() {
			case "string":
				schemaMap.SetPath([]string{k,"default"},one.Get(k).MustString())
				break
			case "integer":
				schemaMap.SetPath([]string{k,"default"},one.Get(k).MustInt())
				break
			case "boolean":
				schemaMap.SetPath([]string{k,"default"},one.Get(k).MustBool())
				break
			}
		}
	 }

	 // 通用信息parent name mid
	 schemaMap.SetPath([]string{"parent","type"},"integer")
	 schemaMap.SetPath([]string{"parent","title"},"上级数据")
	 if nil !=one{
	 	schemaMap.SetPath([]string{"parent","default"},one.Get("parent").MustInt())
	 }


	schemaMap.SetPath([]string{"name","type"},"string")
	schemaMap.SetPath([]string{"name","title"},"名称")
	if nil !=one{
		schemaMap.SetPath([]string{"name","default"},one.Get("name").MustString())
	}


	schemaMap.SetPath([]string{"seq","type"},"integer")
	schemaMap.SetPath([]string{"seq","title"},"排序")
	if nil !=one{
		schemaMap.SetPath([]string{"seq","default"},one.Get("seq").MustInt())
	}


	schemaMap.SetPath([]string{"status","type"},"integer")
	schemaMap.SetPath([]string{"status","title"},"排序")
	if nil !=one{
		schemaMap.SetPath([]string{"status","default"},one.Get("status").MustInt())
	}

	 c.Data["Schema"]=schemaMap.MustMap()
	 //初始化通用form
	 formarrayObj:=formArray.MustArray()//formArr
	 if len(formarrayObj)<=0{
	 	var tmpArray []map[string]string
	 	tmpArray=append(tmpArray, map[string]string{"key":"parent"})
	 	tmpArray=append(tmpArray, map[string]string{"key":"name"})
	 	tmpArray=append(tmpArray, map[string]string{"key":"sql"})
	 	tmpArray=append(tmpArray, map[string]string{"key":"status"})

	 	for k,_:=range schemaMap.MustMap(){
	 		tmpArray=append(tmpArray, map[string]string{"key":k})
		}

	 	tmpArray=append(tmpArray, map[string]string{"type":"submit","title":"提交"})
	 	c.Data["Form"]=tmpArray
	 }else {
	 	var tmpArray []interface{}
		 tmpArray=append(tmpArray, map[string]string{"key":"parent"})
		 tmpArray=append(tmpArray, map[string]string{"key":"name"})
		 tmpArray=append(tmpArray, map[string]string{"key":"sql"})
		 tmpArray=append(tmpArray, map[string]string{"key":"status"})

		 var haveSubmit bool=false
		 for k,v:=range formArray.MustArray(){
		 	tmpArray=append(tmpArray,v)
		 	tmp:=formArray.GetIndex(k).Get("type")
		 	if "submit"==tmp.MustString(){
		 		haveSubmit=true
			}
		 }

		 if false==haveSubmit{
		 	tmpArray=append(tmpArray,map[string]string{"type":"submit","title":"提交"})
		 }
		 c.Data["Form"]=tmpArray
	 }
}

func (c *DataController)EditDo()  {
	did,_:=c.GetInt("did")
	if did>0{
		if len(c.Ctx.Input.RequestBody)>0{
			sj,err:=simplejson.NewJson(c.Ctx.Input.RequestBody)
			if nil ==err{
				var m models.DataModel
				m.Content=string(c.Ctx.Input.RequestBody)
				m.Did=did
				m.Parent=sj.Get("parent").MustInt()
				m.Name=sj.Get("name").MustString()
				m.Mid=c.Mid
				m.Seq=sj.Get("seq").MustInt()
				m.Status= int8(sj.Get("status").MustInt())
				m.UpdateTime=time.Now().Unix()

				id,err:=orm.NewOrm().Update(&m)
				if nil==err{
					c.jsonResult(200,"ok",id)
				}
			}
		}
	}
	c.jsonResult(400,"",0)
}