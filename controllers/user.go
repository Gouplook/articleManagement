package controllers

import (
	"article_management/models"
	"fmt"
	"github.com/beego/beego/v2/adapter/logs"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
)

type UserController struct {
	web.Controller
}

//----注册部分--------

//显示注册页面
func (u *UserController) ShowRegister() {
	u.TplName = "register.html"
}

//处理注册数据
func (u *UserController) HandleRegister() {
	//1.获取数据
	userName := u.GetString("userName")
	pwd := u.GetString("password")

	//2.校验数据
	if userName == "" || pwd == "" {
		u.Data["errmsg"] = "注册数据不完整，请重新注册"
		u.TplName = "register.html"
		return
	}

	//3.操作数据 获取ORM对象
	//获取插入对象 给插入对象赋值
	o := orm.NewOrm()
	var user models.User
	user.Name = userName
	user.PassWord = pwd

	//插入
	_, err := o.Insert(&user)
	if err != nil {
		logs.Info("User Insert fail ...")
		return
	}
	// 返回结果
	// 重定向（浏览器端，不能带数据）
	u.Redirect("/login", 302)

}

//----登陆部分--------

//展示登录页面
func (u *UserController) ShowLogin() {
	u.TplName = "login.html"
}

//处理login数据
func (u *UserController) HandleLogin() {
	//1.获取数据
	userName := u.GetString("userName")
	pwd := u.GetString("password")
	// userName = strings.Trim(userName," ")

	//2.校验数据
	if userName == "" || pwd == "" {
		return
	}

	//3.操作数据(查询）
	user := models.User{
		Name: userName,
	}
	obj := orm.NewOrm()
	err := obj.Read(&user)
	if err == orm.ErrNoRows{
		fmt.Println("查询不到")
		return
	}
	if user.PassWord !=pwd {
		fmt.Println("密码错误")
		return
	}

	fmt.Println("ok----")
	// 返回页面
	u.Redirect("/showArticleList",302)

}


//退出登录
func (u *UserController) Logout() {
	//删除session

	//跳转登录页面
	u.Redirect("/login",302)
}