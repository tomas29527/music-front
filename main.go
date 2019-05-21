package main

import (
	"github.com/astaxie/beego"
	"music-front/models"
	_ "music-front/routers"
)

func main() {
	//日志设置
	logSetting()

	//初始化数据库
	db := &models.Db{}

	if err := db.InitDb() ;err!=nil {
		return
	}

	beego.Run()
}
