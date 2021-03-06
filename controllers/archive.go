package controllers

import (
	"cwengo.com/models"
	"github.com/astaxie/beego"
)

type ArchiveController struct {
	beego.Controller
}

func (this *ArchiveController) Get() {
	this.Data["Categories"], _ = models.GetAllCategories()
	this.Data["Labels"], _ = models.GetAllLabels()
	this.Data["NewTopics"], _ = models.GetAllNewTopics()
	this.Data["VIewsTopics"], _ = models.GetAllViewsTopics()
	this.Data["IsArchive"] = true
	this.Data["ArchiveTopics"], _ = models.GetAllArchiveTopics()
	this.TplNames = "archive.html"
}
