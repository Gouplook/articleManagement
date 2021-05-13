package controllers

import "github.com/beego/beego/v2/server/web"


type UserController struct {
	web.Controller
}


//----注册部分--------

//显示注册页面
func (u *UserController)ShowRegister(){
	u.TplName = "register.html"
}

//处理注册数据
func (u *UserController)HandleRegister(){

}


//----登陆部分--------

//展示登录页面
func (u *UserController)ShowLogin(){
	u.TplName = "login.html"
}


//处理login数据
func (u *UserController)HandleLogin(){

}