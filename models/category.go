package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)

func AddCategory(newCategory string) error {
	o := orm.NewOrm()
	cate := &Category{
		Title:     newCategory,
		Created:   time.Now().Format("2006-01-02 15:04:05"),
		TopicTime: time.Now().Format("2006-01-02 15:04:05"),
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

func EditCategory(category_id, category_name string) error {
	o := orm.NewOrm()
	CateId, err := strconv.ParseInt(category_id, 10, 64)
	if err != nil {
		return err
	}
	cate := &Category{Id: CateId}
	qs := o.QueryTable("category")
	err = qs.Filter("id", category_id).One(cate)
	if err != nil {
		return errors.New("Category id is not exists")
	}
	err = qs.Filter("title", category_name).One(cate)
	if err == nil {
		return errors.New("Category name already exists")
	}
	cate.Title = category_name
	_, err = o.Update(cate)
	return nil
}
