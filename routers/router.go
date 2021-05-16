/**
 * @Author: yinjinlin
 * @File:  router.go
 * @Description:
 * @Date: 2021/5/13 下午1:50
 */

package routers

import (
	"article_management/controllers"
	"github.com/beego/beego/v2/server/web"

)

func init() {
	// 路由
	// 用户模块
	web.Router("/register", &controllers.UserController{}, "get:ShowRegister;post:HandleRegister")
	web.Router("/login", &controllers.UserController{}, "get:ShowLogin;post:HandleLogin")
	//退出登录
	web.Router("/article/logout",&controllers.UserController{},"get:Logout")


	//文章列表页访问
	web.Router("/article/showArticleList",&controllers.AtricleController{},"get:ShowArticleList")
	//添加文章
	web.Router("/article/addArticle",&controllers.AtricleController{},"get:ShowAddArticle;post:HandleAddArticle")
	//显示文章详情
	web.Router("/article/showArticleDetail",&controllers.AtricleController{},"get:ShowArticleDetail")
	//编辑文章
	web.Router("/article/updateArticle",&controllers.AtricleController{},"get:ShowUpdateArticle;post:HandleUpdateArticle")
	//删除文章
	web.Router("/article/deleteArticle",&controllers.AtricleController{},"get:DeleteArticle")
	//添加分类
	web.Router("/article/addType",&controllers.AtricleController{},"get:ShowAddType;post:HandleAddType")
	//删除类型
	web.Router("/article/deleteType",&controllers.AtricleController{},"get:DeleteType")
}
