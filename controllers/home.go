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
	list, _ := models.GetAllTopics("", "", true, start, offset)
	totalCount, _ := models.GetAllTopicsCount()
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
	conf := utils.Config{
		PageUrl:       "/",
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
	this.Data["Categories"], _ = models.GetAllCategories()
	this.Data["PageStr"] = "<ul class='pagination'>" + pageStr + "</ul>"
	this.Data["IsHome"] = true
	this.TplNames = "home.html"
}
