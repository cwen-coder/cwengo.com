package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Admin struct {
	Id       int64
	Username string `orm:"unique"`
	Password string
}

type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}

type Label struct {
	Id              int64
	Title           string    `orm:"index"`
	Created         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}

type Topic struct {
	Id         int64
	Uid        int64
	Title      string
	Category   string
	Lables     string
	Content    string `orm:"site(5000)"`
	Attachment string
	Created    time.Time `orm:"index"`
	Updated    time.Time `orm:"index"`
	Views      int64     `orm:"index"`
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
