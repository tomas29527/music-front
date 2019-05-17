package routers

import (
	"github.com/astaxie/beego"
	"music-front/controllers"
)

func init() {
	beego.Router("/index/bannerList", &controllers.IndexController{}, "get:BannerList")
	beego.Router("/index/recommendList", &controllers.IndexController{}, "get:RecommendList")
}
