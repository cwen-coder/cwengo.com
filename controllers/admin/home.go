package admin

import (
	"cwengo.com/models"
	"cwengo.com/utils"
	"github.com/astaxie/beego"
	"strconv"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Prepare() {
	userid := this.GetSession("userid")
	username := this.GetSession("username")
	if userid == nil || username == nil {
		this.Ctx.Redirect(302, "/admin/login")
		return
	}
}

func (this *HomeController) Get() {

	this.Ctx.Request.ParseForm()

	page, _ := strconv.Atoi(this.Ctx.Request.Form.Get("page"))
	offset := 10
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

	this.Data["Topics"] = list
	//分页配置
	conf := utils.Config{
		PageUrl:       "/admin/home",
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
		this.Ctx.Redirect(302, "/admin/home")
	}
	this.Data["PageStr"] = "<ul class='pagination'>" + pageStr + "</ul>"
	this.Data["Username"] = this.GetSession("username")
	this.Data["toplicList"] = true
	this.Layout = "admin/layout.html"
	this.TplNames = "admin/Tpl/T.home.tpl"
}
