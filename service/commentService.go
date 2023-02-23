package service

import (
	"douyinOrigin/dao"
)

// CommentInfo 查看评论-传出的结构体-service
type CommentInfo struct {
	Content    string `json:"content"`     // 评论内容
	CreateDate string `json:"create_date"` // 评论发布日期，格式 mm-dd
	ID         int64  `json:"id"`          // 评论id
	User       User   `json:"user"`        // 评论用户信息
}

// CommentService 接口定义
// 发表评论-使用的结构体-service层引用dao层↑的Comment。
type CommentService interface {

	// CountFromVideoId 1.根据videoId获取视频评论数量的接口
	CountFromVideoId(id int64) (int64, error)

	// Send 2、发表评论，传进来评论的基本信息，返回保存是否成功的状态描述
	Send(comment dao.Comment) (CommentInfo, error)

	// DelComment 3、删除评论，传入评论id即可，返回错误状态信息
	DelComment(commentId int64) error

	// GetList 4、查看评论列表-返回评论list-在controller层再封装外层的状态信息
	GetList(videoId int64, userId int64) ([]CommentInfo, error)
}
