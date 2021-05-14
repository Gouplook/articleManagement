/**
 * @Author: yinjinlin
 * @File:  db
 * @Description:
 * @Date: 2021/5/13 下午2:22
 */

package models

import (
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init(){
	//ORM操作数据库
	//获取连接对象
	_ = orm.RegisterDataBase("default","mysql","root:123456@tcp(127.0.0.1:3306)/articlemanagement?charset=utf8")

	//register model
	orm.RegisterModel(new(User),new(Article),new(ArticleType))

	// register model 表名前缀
	// orm.RegisterModelWithPrefix("mac_", new(User),new(Article),new(ArticleType))

	//生成表
	//第一个参数是数据库别名，第二个参数是是否强制更新(多表操作时需要强制更新）
	_ = orm.RunSyncdb("default", true, true)

	// 设置数据库的最大空闲连接
	orm.SetMaxIdleConns("default", 30)

	// 设置数据库的最大数据库连接
	orm.SetMaxOpenConns("default", 30)


}