package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)

func AddCategory(newCategory string) error {
	o := orm.NewOrm()
	cate := &Category{
		Title:     newCategory,
		Created:   time.Now(),
		TopicTime: time.Now(),
	}
	qs := o.QueryTable("category")
	err := qs.Filter("title", newCategory).One(cate)
	if err == nil {
		return err
	}
	_, err = o.Insert(cate)

	if err != nil {
		return err
	}
	return nil
}

func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()
	cates := make([]*Category, 0)
	qs := o.QueryTable("category")
	_, err := qs.All(&cates)
	return cates, err
}

func DelCategory(category_id string) error {
	o := orm.NewOrm()
	CateId, err := strconv.ParseInt(category_id, 10, 64)
	if err != nil {
		return err
	}
	cate := &Category{Id: CateId}
	_, err = o.Delete(cate)
	return err
}
