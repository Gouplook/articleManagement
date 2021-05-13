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

}
