package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"music-front/models"
	"music-front/utils"
	"music-front/vo"
)

//报名

type SignUpController struct {
	beego.Controller
}

func (s *SignUpController) SignUp() {
	var result *vo.Result
	signUp := &vo.SignUpVo{}
	logs.Info("=======:%v", s.Input())
	if err := s.ParseForm(signUp); err != nil {
		logs.Error("解析参数异常:%v", err)
		result = vo.NewResultFail("参数错误")
		s.Data["json"] = result
		s.ServeJSON()
		return
	}
	logs.Info("请求参数:SignUp=:%v", signUp)
	params, e := utils.ValidParams(signUp)
	if e != nil {
		logs.Error("解析参数异常:%v", e)
		result = vo.NewResultFail("参数错误")
		s.Data["json"] = result
		s.ServeJSON()
		return
	}
	if params != "" {
		result = vo.NewResultFail(params)
	}
	up := models.NewSignUp(signUp)
	er := up.InsertSignUp()
	if er != nil {
		logs.Error("保存出错:%v", er)
		result = vo.NewResultFail("参数错误")
		s.Data["json"] = result
		s.ServeJSON()
		return
	}
	result = vo.NewResultSuccess()
	s.Data["json"] = result
	s.ServeJSON()
}
