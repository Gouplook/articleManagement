/**
 * @Author: yinjinlin
 * @File:  main.go
 * @Description:
 * @Date: 2021/5/13 下午1:46
 */

package main

import (
	_ "article_management/models" // 初始化models下 init
	_ "article_management/routers"
	"github.com/beego/beego/v2/server/web"
)

func main() {
	// 视图函数映射
	_ = web.AddFuncMap("pagepre",ShowPrePage)
	_ = web.AddFuncMap("nextpage",ShowNextPage)


	web.Run()
}

// 上一页视图函数（index.html）
func ShowPrePage(pageIndex int) int {
	if pageIndex == 1 {
		return pageIndex
	}
	return pageIndex - 1
}

// 下一页视图函数（index.html）
func ShowNextPage(pageIndex int, pageCount int) int {
	if pageIndex == pageCount {
		return pageCount
	}
	return pageIndex + 1
}