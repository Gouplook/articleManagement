/**
 * @Author: yinjinlin
 * @File:  userModel
 * @Description:
 * @Date: 2021/5/13 下午2:22
 */

package models

// 表设计关系
// 用户    <--->  文章   m2m
// 文章类型 ---->  文章  （一对多）

// 用户表
type User struct {
	Id int
	Name string
	PassWord string
	//Pass_Word
	//Articles []*Article `orm:"reverse(many)"`

}
