package controllers

import (
	"cwengo.com/models"
	"cwengo.com/utils"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.Request.ParseForm()
	page, _ := strconv.Atoi(this.Ctx.Request.Form.Get("page"))
	offset := 8
	if page == 0 {
		page = 1
	}

	start := (page - 1) * offset
	cate := this.Input().Get("cate")
	label := this.Input().Get("label")
	list, _ := models.GetAllTopics(cate, label, true, start, offset)
	totalCount, _ := models.GetAllTopicsCount(cate, label)
	var pageCount int
	if totalCount%offset == 0 {
		pageCount = totalCount / offset
	} else {
		pageCount = totalCount/offset + 1
	}
	LabelId := make([][]string, 10)
	Label := make([][]string, 10)
	for i, v := range list {
		LabelId[i] = strings.Split(v.Lables, " ")
		LabelLen := len(LabelId[i])
		LabelId[i] = LabelId[i][0 : LabelLen-1]
		LabelOne := make([]string, 10)
		for j, n := 0, len(LabelId[i]); j < n; j++ {
			LabelOne[j] = models.GetLabel(LabelId[i][j])
		}
		Label[i] = LabelOne
	}
	this.Data["Topics"] = list
	this.Data["LabelId"] = LabelId
	this.Data["Label"] = Label
	//分页配置
	var PageUrl string
	if cate != "" {
		PageUrl = "/?cate=" + cate
		this.Data["HomeTitle"] = "分类：" + models.GetCategory(cate)
	} else if label != "" {
		PageUrl = "/?label=" + label
		this.Data["HomeTitle"] = "标签：" + models.GetLabel(label)
	} else {
		PageUrl = "/"
		this.Data["HomeTitle"] = "最新文章"
	}
	conf := utils.Config{
		PageUrl:       PageUrl,
		PageSize:      1,
		RowsNum:       pageCount,
		AnchorClass:   "",
		LinksNum:      1,
		PageNum:       page,
		CurrentClass:  "",
		InfoTagOpen:   "<li>",
		InfoTagClose:  "</li>",
		FirstTagOpen:  "<li>",
		FirstTagClose: "</li>",
		LastTagOpen:   "<li>",
		LastTagClose:  "</li>",
		CurTagOpen:    "<li>",
		CurTagClose:   "</li>",
		NextTagOpen:   "<li>",
		NextTagClose:  "</li>",
		PrevTagOpen:   "<li>",
		PrevTagClose:  "</li>",
		NumTagOpen:    "<li>",
		NumTagClose:   "</li>",
	}

	pageStr, err := utils.CreateLinks(conf)
	if err != nil {
		beego.BeeLogger.Critical("CreateLinks: ", err)
	}

	if pageStr == "404" {
		this.Ctx.Redirect(302, "/")
	}

	this.Data["PageStr"] = "<ul class='pagination'>" + pageStr + "</ul>"
	this.Data["Categories"], _ = models.GetAllCategories()
	this.Data["Labels"], err = models.GetAllLabels()
	this.Data["NewTopics"], err = models.GetAllNewTopics()
	this.Data["VIewsTopics"], err = models.GetAllViewsTopics()
	this.Data["IsHome"] = true
	this.TplNames = "home.html"
}
