package models

import (
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
	LabelTime       time.Time `orm:"index"`
	LabelCount      int64
	LabelLastUserId int64
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

func RegisterDB() {
	orm.RegisterModel(new(Admin), new(Category), new(Topic), new(Label))
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", "root:yin123@/cwengo.com?charset=utf8")
}
