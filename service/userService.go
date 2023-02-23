package service

import (
	"douyinOrigin/dao"
)

// User 最终封装后,controller返回的User结构体
type User struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	FollowCount   int    `json:"follow_count"`
	FollowerCount int    `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`

	FavoriteCount   int64  `json:"favorite_count"`   // 喜欢数
	TotalFavorited  string `json:"total_favorited"`  // 获赞数量
	Signature       string `json:"signature"`        // 个人简介
	WorkCount       int64  `json:"work_count"`       // 作品数
	Avatar          string `json:"avatar"`           // 用户头像
	BackgroundImage string `json:"background_image"` // 用户个人页顶部大图
}

//type User struct {
//	FavoriteCount  int64  `json:"favorite_count"`  // 喜欢数
//	FollowCount    int64  `json:"follow_count"`    // 关注总数
//	FollowerCount  int64  `json:"follower_count"`  // 粉丝总数
//	Id             int64  `json:"id"`              // 用户id
//	IsFollow       bool   `json:"is_follow"`       // true-已关注，false-未关注
//	Name           string `json:"name"`            // 用户名称
//	TotalFavorited string `json:"total_favorited"` // 获赞数量
//
//	Avatar          string `json:"avatar"`           // 用户头像
//	BackgroundImage string `json:"background_image"` // 用户个人页顶部大图
//	Signature       string `json:"signature"`        // 个人简介
//	WorkCount       int64  `json:"work_count"`       // 作品数
//}

type UserService interface {
	// GetTableUserList 获取TableUser对象
	GetTableUserList() []dao.TableUser

	// GetTableUserByUsername 根据user_name获得User对象
	GetTableUserByUsername(name string) dao.TableUser

	// GetTableUserById 根据 user_id 获得User对象
	GetTableUserById(id int64) dao.TableUser

	// InsertTableUser 将user插入到数据表中
	InsertTableUser(tableUser *dao.TableUser) bool

	// GetUserByCurId 在登录情况（curId）下根据user_id获取User对象，curId判断是否点赞
	GetUserByCurId(id int64, curId int64) (User, error)
}
