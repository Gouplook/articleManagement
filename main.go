/**
 * @Author: yinjinlin
 * @File:  main.go
 * @Description:
 * @Date: 2021/5/13 下午1:46
 */

package main

import (
	_ "article_management/models"   // 初始化models下 init
	_ "article_management/routers"
	"github.com/beego/beego/v2/server/web"
)

func main() {

	web.Run()
}
