package admin

import (
	/*"cwengo.com/models"*/
	/*"cwengo.com/utils"*/
	"github.com/astaxie/beego"
	/*"strconv"*/)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	/*	this.Ctx.Request.ParseForm()

		page, _ := strconv.Atoi(this.Ctx.Request.Form.Get("page"))
		offset := 1
		if page == 0 {
			page = 1
		}

		start := (page - 1) * offset
		list := models.GetBlogsList(start, offset)
		countInfo, _ := models.GetAllBlogsCount()
		count := string(countInfo[0]["count(*)"])
		totalCount, _ := strconv.Atoi(count)
		newList := make([]map[string]interface{}, len(list))
		for k, v := range list {
			m := make(map[string]interface{})
			m["id"] = v.Id
			m["title"] = v.Title
			m["content"] = v.Content
			m["created"] = v.Created
			newList[k] = m
		}

		this.Data["Blogs"] = newList
		//分页配置
		conf := utils.Config{
			PageUrl:       "/admin/index",
			PageSize:      1,
			RowsNum:       totalCount,
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
			this.Ctx.Redirect(302, "/admin/index")
		}
		this.Data["PageStr"] = "<ul>" + pageStr + "</ul>"*/
	this.Data["Username"] = this.GetSession("username")
	this.Data["toplicList"] = true
	this.Layout = "admin/layout.html"
	this.TplNames = "admin/Tpl/T.home.tpl"
}
