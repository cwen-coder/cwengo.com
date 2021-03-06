package models

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)

func AddLabel(newLabel string) error {
	o := orm.NewOrm()
	label := &Label{
		Title:     newLabel,
		Created:   time.Now().Format("2006-01-02 15:04:05"),
		TopicTime: time.Now().Format("2006-01-02 15:04:05"),
	}
	qs := o.QueryTable("label")
	err := qs.Filter("title", newLabel).One(label)
	if err == nil {
		return err
	}
	_, err = o.Insert(label)

	if err != nil {
		return err
	}
	return nil
}

func GetAllLabels() ([]*Label, error) {
	o := orm.NewOrm()
	Labels := make([]*Label, 0)
	qs := o.QueryTable("label")
	_, err := qs.All(&Labels)
	return Labels, err
}

func DelLabel(label_id string) error {
	o := orm.NewOrm()
	LabelId, err := strconv.ParseInt(label_id, 10, 64)
	if err != nil {
		return err
	}
	label := &Label{Id: LabelId}
	_, err = o.Delete(label)
	return err
}

func EditLabel(label_id, label_name string) error {
	o := orm.NewOrm()
	LabelId, err := strconv.ParseInt(label_id, 10, 64)
	if err != nil {
		return err
	}
	label := &Label{Id: LabelId}
	err = o.Read(label)
	if err != nil {
		return errors.New("Label id is not exists")
	}
	oldLabel := &Label{Title: label_name}
	err = o.Read(oldLabel, "Title")
	if err == nil {
		return errors.New("Label name already exists")
	}
	label.Title = label_name
	_, err = o.Update(label)
	return nil
}

func GetLabel(id string) string {
	o := orm.NewOrm()
	LabelId, err := strconv.ParseInt(id, 10, 64)
	label := &Label{Id: LabelId}
	qs := o.QueryTable("label")
	err = qs.Filter("id", LabelId).One(label)
	if err != nil {
		beego.Error(err)
	}
	return label.Title
}

func GetLabel1(id string) *Label {
	o := orm.NewOrm()
	LabelId, err := strconv.ParseInt(id, 10, 64)
	label := &Label{Id: LabelId}
	qs := o.QueryTable("label")
	err = qs.Filter("id", LabelId).One(label)
	if err != nil {
		beego.Error(err)
	}
	return label
}
