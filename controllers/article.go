package controllers

import (
	"article_management/models"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
)

type AtricleController struct {
	web.Controller
}


// 展示文章列表页
func (a *AtricleController)ShowArticleList(){
	//session判断

	// 获取数据（高级查询）
	//obj := orm.NewOrm()
	//qs := obj.QueryTable("Article")

	//查询总记录数

	//分页处理
	//获取总页数
}

//展示添加文章页面
func (a *AtricleController)ShowAddArticle(){
	//查询所有类型数据，并展示
	o := orm.NewOrm()
	var types []models.ArticleType
	o.QueryTable("ArticleType").All(&types)

	//传递数据(查询的数据传提给页面）
	a.Data["types"] = types
	a.TplName = "add.html"

}


//获取添加文章数据 (post)
func (a *AtricleController)HandleAddArticle(){
	//1.获取数据

	//2校验数据

	//处理文件上传(本地保存)
	file,fileHear,err :=  a.GetFile("uploadname")
	defer  file.Close()

	if err != nil {
		return
	}


	//1.文件大小
	if fileHear.Size > 5000000{

	}

	//2.文件格式
	//ext := path.Ext(fileHear.Filename)

	//3.防止重名
	fileName := ""

	//存储
	//a.SaveToFile()

	//3.处理数据(数据插入数据库）
	var artice models.Article
	artice.Aimg = "static/img"+fileName

	//给文章添加类型
	//获取类型数据
	typeName := ""

	//根据名称查询类型
	var articType models.ArticleType
	articType.TypeName = typeName


	//4.返回页面

}

//展示文章详情页面
func (a *AtricleController)ShowArticleDetail(){
	// 	获取数据

	// 数据校验


	// 查询数据
	//obj := orm.NewOrm()


	// 修改阅读量 *** update

	//多对多 插入浏览记录 ****
	//1 获取orm对象
	//2 获取操作对象
	//3 获取对多对操作对象
	//4 获取插入对象
	//5 插入

	//obj.QueryM2M()


	//返回视图页面

	a.Layout = "layout.html"
	a.TplName = "index.html"

}

//显示编辑页面
func (a *AtricleController)ShowUpdateArticle()  {

	// 	获取数据

	// 数据校验


	// 查询数据

	//返回视图页面
	a.TplName = "update.html"

}


//封装上传文件函数...
func upLoadFile(this *web.Controller,filePath string) string{

	return ""
}



//处理编辑界面数据
func(a *AtricleController)HandleUpdateArticle(){

	// 获取数据(页面数据）

	//中间涉及图片上传的数据问题
	//upLoadFile()

	//校验数据(页面数据）

	//数据更新update

	//返回视图 **
	a.Redirect("/article/showArticleList",302)
}

//删除文章处理
func (a *AtricleController)DeleteArticle() {
	// 获取需要删除的数据

	//数据操作delete

	//返回视图
	a.Redirect("xxxx",302)
}


//展示添加类型页面
func (a *AtricleController)ShowAddType(){
	//查询数据 atricleType

	// 传提数据到页面
	a.TplName ="xxx.html"

}

//处理添加类型数据
func (a *AtricleController)HandleAddType()  {

	// 获取数据（页面的数据）

	//数据插入insert

	//返回视图
	a.Redirect("xxxx", 302)
}

//删除类型
func (a *AtricleController)DeleteType() {
	// 获取需要删除的数据

	//执行删除

	//返回视图(添加的类型页面）
	a.Redirect("xxxx", 302)
}
