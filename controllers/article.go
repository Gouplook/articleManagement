package controllers

import (
	"article_management/models"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	"math"
	"path"
	"strings"
	"time"
)

type AtricleController struct {
	web.Controller
}

// 展示文章列表访问页
func (a *AtricleController) ShowArticleList() {
	// session判断

	// 获取数据（高级查询）
	obj := orm.NewOrm()
	var articles []models.Article
	id,_ := obj.QueryTable("Article").All(&articles)
	if id < 0 {
		// 需要收集日志...
		return
	}

	// 每页显示的多少条
	pageSize := 2
	// （页码）获取前端传过来的页码
	pageIndex,err := a.GetInt("xxxx")
	if err != nil {
		pageIndex = 1
	}
	//起始位置计算
	start := ( pageIndex - 1) * pageSize
	typeName := a.GetString("select")
	var count int64
	if typeName == "" {
		// 查询总记录数
		count, _ = obj.QueryTable("Article").Count()
	}else {
		count,_ = obj.QueryTable("Article").
			Limit(pageSize,start).
			RelatedSel("ArticleType").
			Filter("ArticleType__TypeName",typeName).Count()
	}
	// 获取总页数
	pageCount := math.Ceil(float64(count) / float64(pageSize))

	// 获取文章类型
	var types []models.ArticleType
	id,_ = obj.QueryTable("ArticleType").All(&types)
	if id < 0 {
		// 需要收集日志...
		return
	}

	if typeName != "" {
		id,_ = obj.QueryTable("Article").
			Limit(pageSize,start).
			RelatedSel("ArticleType").
			Filter("ArticleType__TypeName",typeName).
			All(&articles)
	}
	if id < 0 {
		// 需要收集日志...
		return
	}



	// 需要传送数据到页面
	a.Data["articles"] = articles
	a.Data["count"] = count
	a.Data["types"] = types
	a.Data["pageCount"] = int(pageCount)
	a.Data["pageIndex"] = pageIndex
	a.Data["typeName"] = typeName

	// 返回视图
	a.TplName = "index.html"
}

// 展示添加文章页面
func (a *AtricleController) ShowAddArticle() {
	// 查询所有类型数据，并展示
	obj := orm.NewOrm()
	var types []models.ArticleType
	id,_ := obj.QueryTable("ArticleType").All(&types)
	if id < 0 {
		// 需要收集日志...
		return
	}
	// 传递数据(查询的数据传提给页面）
	a.Data["types"] = types
	a.TplName = "add.html"

}

// 添加文章数据 (post)
func (a *AtricleController) HandleAddArticle() {
	// 1.获取数据
	articleName := a.GetString("articleName")
	articleContent := a.GetString("content")

	articleName = strings.Trim(articleName, "")
	articleContent = strings.Trim(articleContent, "")

	// 2校验数据
	if articleName == "" || articleContent == "" {
		a.Data["errMsg"] = "添加的文章的内容不用为空"
		a.Redirect("add.html", 302)
		// return
	}

	// 处理文件上传(本地保存)
	file, fileHear, err := a.GetFile("uploadName")
	defer file.Close()
	if err != nil {
		a.Data["errMsg"] = "上传的照片有误，请重新选择"
		a.Redirect("add.html", 302)
	}

	// 1.文件大小
	if fileHear.Size > 5000000 {
		a.Data["errMsg"] = "上传的照片过大，请重新选择"
		// a.Redirect("add.html",302) 不能使用重定向，由于需要带上数据
		a.TplName = "add.html"
	}

	// 2.文件格式 获取文件的后缀
	ext := path.Ext(fileHear.Filename)
	if ext != ".jpg" && ext != ".png" && ext != ".jpeg" {
		a.Data["errMsg"] = "文件格式错误。请重新上传"
		a.TplName = "add.html"
	}

	// 3.防止重名
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	fileName := timeStr + ext

	// 存储
	// 第一个参数 页面的input 的name
	_ = a.SaveToFile("uploadName", ".static/img/"+fileName)

	// 3.处理数据(数据插入数据库）
	var article models.Article
	article.ArtiName = articleName
	article.Acontent = articleContent
	article.Aimg = ".static/img/" + fileName
	article.Atime = time.Now()

	// 给文章添加类型 获取类型数据
	typeName := a.GetString("select")
	// 根据名称查询类型
	var articType models.ArticleType
	articType.TypeName = typeName
	obj := orm.NewOrm()
	err = obj.Read(&articType, "TypeName")
	if err != nil {
		// 需要收集日志
		return
	}
	// 后续 ....
	// artice.ArticleType = &articType

	// 添加的文章表插入
	id, _ := obj.Insert(&article)
	if id < 0 {
		// 需要收集日志
		return
	}

	// 4.返回页面
	a.Redirect("/ShowArticleList", 302)
}

// 展示文章详情页面
func (a *AtricleController) ShowArticleDetail() {
	// 	获取数据

	// 数据校验

	// 查询数据
	// obj := orm.NewOrm()

	// 修改阅读量 *** update

	// 多对多 插入浏览记录 ****
	// 1 获取orm对象
	// 2 获取操作对象
	// 3 获取对多对操作对象
	// 4 获取插入对象
	// 5 插入

	// obj.QueryM2M()

	// 返回视图页面

	a.Layout = "layout.html"
	a.TplName = "index.html"

}

// 显示编辑页面
func (a *AtricleController) ShowUpdateArticle() {

	// 	获取数据

	// 数据校验

	// 查询数据

	// 返回视图页面
	a.TplName = "update.html"

}

// 封装上传文件函数...
func upLoadFile(this *web.Controller, filePath string) string {

	return ""
}

// 处理编辑界面数据
func (a *AtricleController) HandleUpdateArticle() {

	// 获取数据(页面数据）

	// 中间涉及图片上传的数据问题
	// upLoadFile()

	// 校验数据(页面数据）

	// 数据更新update

	// 返回视图 **
	a.Redirect("/article/showArticleList", 302)
}

// 删除文章处理
func (a *AtricleController) DeleteArticle() {
	// 获取需要删除的数据

	// 数据操作delete

	// 返回视图
	a.Redirect("xxxx", 302)
}

// 获取添加类型页面
func (a *AtricleController) ShowAddType() {
	// 使用高级查询 atricleType
	obj := orm.NewOrm()
	var types []models.ArticleType
	id, _ := obj.QueryTable("ArticleType").All(&types)

	if id < 0 {
		// 需要收集日志...
		return
	}

	// 传递数据到页面
	a.Data["types"] = types
	a.TplName = "addType.html"
}

// 添加类型数据
func (a *AtricleController) HandleAddType() {
	// 获取数据（页面的数据）
	name := a.GetString("typeName")

	if name == "" {
		a.Data["errMsg"] = "信息不完整，请重新输入"
		return
	}

	var artileType models.ArticleType
	artileType.TypeName = name

	// 插入数据之前对数据进行查找
	obj := orm.NewOrm()

	// 数据插入insert
	id, _ := obj.Insert(&artileType)
	if id < 0 {
		// 需要收集日志...
		return
	}

	// 返回视图 /article/addArticle
	a.Redirect("/article/addArticle", 302)
}

// 删除类型
func (a *AtricleController) DeleteType() {
	// 获取需要删除的数据

	// 执行删除

	// 返回视图(添加的类型页面）
	a.Redirect("xxxx", 302)
}
