package userService

import "douyinOrigin/dao/userDao"

// User 最终封装后,controller返回的User结构体
type User struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	FollowCount   int    `json:"follow_count"`
	FollowerCount int    `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

type UserService interface {
	// GetTableUserList 获取TableUser对象
	GetTableUserList() []userDao.TableUser

	// GetTableUserByUsername 根据user_name获得User对象
	GetTableUserByUsername(name string) userDao.TableUser

	// GetTableUserById 根据 user_id 获得User对象
	GetTableUserById(id int64) userDao.TableUser

	// InsertTableUser 将user插入到数据表中
	InsertTableUser(tableUser *userDao.TableUser) bool
}
