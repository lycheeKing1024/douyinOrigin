package dao

import (
	"fmt"
)

//import "log"

// TableUser 定义User模型，绑定users表，ORM库操作数据库， 需要定义一个struct类型和MYSQL表进行绑定或者叫映射，
// struct字段和MYSQL表字段一一对应
type TableUser struct {
	Id       int64 //主键
	Name     string
	Password string
	//NickName string
	Avatar          string // 用户头像
	BackgroundImage string // 用户个人页顶部大图
	Signature       string // 个人简介
	WorkCount       int64  // 作品数
}

// TableName 设置表名，可以通过struct类型定义 TableName函数，返回当前struct绑定的mysql表名是什么
func (u TableUser) TableName() string {
	//绑定MySQL表名 users
	return "users"
}

// GetTableUserList 获取全部的User对象
func GetTableUserList() ([]TableUser, error) {
	TableUsers := []TableUser{}
	if err := SqlSession.Find(&TableUsers).Error; err != nil {
		//log.Panicln(err.Error())
		fmt.Println(err.Error())
		return TableUsers, err //将查询结果和错误信息返回
	}
	return TableUsers, nil
}

// GetTableUserByUsername 根据user_name获得User对象
func GetTableUserByUsername(username string) (TableUser, error) {
	tableUser := TableUser{}
	if err := SqlSession.Where("name=?", username).First(&tableUser).Error; err != nil {
		//log.Panicln(err.Error())
		//fmt.Println(err.Error())
		return tableUser, err
	}
	return tableUser, nil
}

// GetTableUserById 根据 user_id 获得User对象
func GetTableUserById(id int64) (TableUser, error) {
	tableUser := TableUser{}
	if err := SqlSession.Where("id=?", id).First(&tableUser).Error; err != nil {
		//log.Panicln(err.Error())
		fmt.Println(err.Error())
		return tableUser, err
	}
	return tableUser, nil
}

// InsertTableUser 将user插入到数据表中
func InsertTableUser(user *TableUser) bool {
	if err := SqlSession.Create(&user).Error; err != nil {
		//log.Panicln(err.Error())
		fmt.Println(err.Error())
		return false
	}
	return true
}
