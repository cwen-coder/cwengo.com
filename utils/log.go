package utils

import (
	"github.com/astaxie/beego"
	"time"
)

func init() {
	//获取当天时间
	now := time.Now().Format("2006-01-02")
	fn := "log/" + now + ".log"
	beego.SetLogger("file", `{"filename":"`+fn+`"}`)
}

func Info(v ...interface{}) {
	beego.Info(v)
}
