/**
 * @Author: yinjinlin
 * @File:  articleTypeModel
 * @Description:
 * @Date: 2021/5/14 上午11:23
 */

package models

// 文章类型
type ArticleType struct {
	Id int
	TypeName string `orm:"size(20)"`

	Articles []*Article `orm:"reverse(many)"`

}
