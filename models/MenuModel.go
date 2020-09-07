package models

type MenuModel struct {
	Mid int `orm:"pk;auto"`
	Parent int
	Name string`orm:"size(45)"`
	Seq int
	Format string `orm:"size(2048)";default({})`
}



type MenuTree struct {
	MenuModel
	Child[]MenuModel
}

func (m *MenuModel)TableName()string  {
	return "xcms_menu"
}