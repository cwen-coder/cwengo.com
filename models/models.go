package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Admin struct {
	Id       int64
	Username string `orm:"unique"`
	Password string
}

type Category struct {
	Id              int64
	Title           string
	Created         string `orm:"index"`
	Views           int64  `orm:"index"`
	TopicTime       string `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
	IsSelected      bool
}

type Label struct {
	Id              int64
	Title           string `orm:"index"`
	Created         string `orm:"index"`
	Views           int64  `orm:"index"`
	TopicTime       string `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
	IsSelected      bool
}

type Topic struct {
	Id         int64
	Uid        int64
	Title      string
	Category   string
	Lables     string
	Summery    string `orm:"site(1000)"`
	Content    string `orm:"site(6000)"`
	Attachment string
	Created    string `orm:"index"`
	Updated    string `orm:"index"`
	Views      int64  `orm:"index"`
	Author     string
}

type TopicLabel struct {
	Id      int64
	TopicId int64 `orm:"index"`
	LabelId int64 `orm:"index"`
}

func RegisterDB() {
	orm.RegisterModel(new(Admin), new(Category), new(Topic), new(Label), new(TopicLabel))
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", "root:yin123@/cwengo.com?charset=utf8")
}

func GetAdminInfo(username string) (admin Admin) {
	o := orm.NewOrm()
	qs := o.QueryTable("admin")
	err := qs.Filter("username", username).One(&admin)
	if err != nil {
		fmt.Println("Fail error", err.Error())
	}
	return
}
