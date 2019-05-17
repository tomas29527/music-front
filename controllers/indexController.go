package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"music-front/models"
	"music-front/vo"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) BannerList() {
	bannerList, err := models.QueryBannerList()
	var result *vo.Result
	if err != nil {
		logs.Error("查询QueryBannerList出错:%v", err)
		result = vo.NewResultFail("查询错误")
	} else {
		result = vo.NewResultSuccess(bannerList)
	}
	c.Data["json"] = result
	c.ServeJSON()
}

func (c *IndexController) RecommendList() {
	recommendList, err := models.QueryRecommendList()
	var result *vo.Result
	if err != nil {
		logs.Error("查询QueryRecommendList出错:%v", err)
		result = vo.NewResultFail("查询错误")
	} else {
		result = vo.NewResultSuccess(recommendList)
	}
	c.Data["json"] = result
	c.ServeJSON()
}
