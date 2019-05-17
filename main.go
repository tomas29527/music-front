package main

import (
	"github.com/astaxie/beego"
	"music-front/models"
	_ "music-front/routers"
)

func main() {
	db := &models.Db{}
	defer db.Close()
	//日志设置
	logSetting()

	beego.Run()
}
