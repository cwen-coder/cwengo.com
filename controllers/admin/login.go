package admin

import (
	"bytes"
	"crypto/md5"
	"cwengo.com/models"
	"fmt"
	"github.com/astaxie/beego"
	"io"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplNames = "admin/signin.html"
}

func (this *LoginController) Post() {
	this.TplNames = "admin/signin.html"
	username := this.Input().Get("uname")
	pwd := this.Input().Get("pwd")

	if username == "" {
		this.Data["UsernameNull"] = "username is not null"
		return
	}

	if pwd == "" {
		this.Data["PasswordNull"] = "password is not null"
		return
	}

	adminInfo := models.GetAdminInfo(username)

	if adminInfo.Username == "" {
		this.Data["UsernameNull"] = "用户不存在"
		return
	}

	md5Password := md5.New()
	io.WriteString(md5Password, pwd)
	buffer := bytes.NewBuffer(nil)
	fmt.Fprintf(buffer, "%x", md5Password.Sum(nil))
	newPass := buffer.String()
	if adminInfo.Password != newPass {
		this.Data["PasswordNull"] = "密码错误"
		return
	}
	this.SetSession("userid", adminInfo.Id)
	this.SetSession("username", adminInfo.Username)
	this.Ctx.Redirect(302, "/admin/home")
}
