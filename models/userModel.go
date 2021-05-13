/**
 * @Author: yinjinlin
 * @File:  userModel
 * @Description:
 * @Date: 2021/5/13 下午2:22
 */

package models

type User struct {
	Id int
	Name string
	PassWord string
	//Pass_Word
	//Articles []*Article `orm:"reverse(many)"`
}


