/**
 * @Author: yinjinlin
 * @File:  articleModel
 * @Description:
 * @Date: 2021/5/14 上午11:19
 */

package models

import "time"

// 文章类型 ---->  文章  （一对多）

// 文章
type Article struct {
	Id       int       `orm:"pk;auto"`
	ArtiName string    `orm:"size(20)"`
	Atime    time.Time `orm:"auto_now;type(datetime)"`
	Acount   int       `orm:"default(0);null"`
	Acontent string    `orm:"size(500);description(文章内容)"`
	Aimg     string    `orm:"size(100)"`

	ArticleType *ArticleType `orm:"rel(fk);on_delete(set_null);null"`

	// ArticleType *ArticleType `orm:"rel(fk)"`
	// Users       []*User      `orm:"rel(m2m);rel_table(the_table_name)"`
	// Users       []*User      `orm:"rel(m2m)"`
}
