package VideoService

import (
	"douyinOrigin/dao/videoDao"
	"douyinOrigin/service/userService"
	"mime/multipart"
	"time"
)

type Video struct {
	videoDao.TableVideo
	Author        userService.User
	FavoriteCount int  `json:"favorite_count"`
	CommentCount  int  `json:"comment_count"`
	IsFavorite    bool `json:"is_favorite"`
}
type VideoService interface {
	// Feed 通过传入时间戳，当前用户的id，返回对应的视频切片，以及视频最早的发布时间
	Feed(lastTime time.Time, userId int64) ([]Video, time.Time, error)

	//	传入视频id获得对应的视频对象
	GetVideo(VideoId int64, userId int64) (Video, error)

	// PublishVideo 将传入的视频流保存到文件服务器中，并存储在mysql中
	PublishVideo(data *multipart.FileHeader, userId int64, title string) error

	// TableVideo 通过userId查询对应用户发布的视频，并返回视频数组
	TableVideo(userId int64, curID int64) ([]Video, error)

	// GetVideoIdList 通过一个作者id，返回用户发布的视频id切片数组
	GetVideoIdList(userId int64) ([]int64, error)
}
