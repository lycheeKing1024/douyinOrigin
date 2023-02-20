package videoDao

import (
	"douyinOrigin/config"
	"douyinOrigin/dao"
	"time"
)

type VideoList struct {
	Id       int64  `json:"id"` //视频id
	AuthorId int64  //视频作者信息
	PlayUrl  string `json:"play_url"`  //视频播放地址
	CoverUrl string `json:"cover_url"` //视频封面地址
	//FavoriteCount int    `json:"favorite_count"`
	//CommentCount  int    `json:"comment_count"`
	//IsFavorite    bool   `json:"is_favorite"`
	PublishTime time.Time
	Title       string `json:"title"` //视频标题
}

// TableName 将 VideoList 结构体映射到数据表 videos
func (VideoList) TableName() string {
	return "videos"
}

// GetVideoByVeidoId 根据视频id获取视频信息
func GetVideoByVeidoId(VideoId int64) (VideoList, error) {
	var VideoList VideoList
	VideoList.Id = VideoId
	result := dao.SqlSession.First(&VideoList)
	if result.Error != nil {
		return VideoList, result.Error
	}
	return VideoList, nil
}

// GetVideoByAuthorId 根据作者id查询视频信息，返回切片
func GetVideoByAuthorId(authorId int64) ([]VideoList, error) {
	var VideoLists []VideoList
	result := dao.SqlSession.Where("authorId=?", authorId).Find(&VideoLists)
	if result.Error != nil {
		//如果出现问题，返回对应到空，并且返回错误信息
		return nil, result.Error
	}
	return VideoLists, nil
}

// GetVideoByLastTime 根据传入的时间来获取这个时间之前的一些数据
func GetVideoByLastTime(lastTime time.Time) ([]VideoList, error) {
	Videos := make([]VideoList, 0, config.VideoNum)
	result := dao.SqlSession.Where("publish_time<?", lastTime).Order("publish_time desc").Limit(config.VideoNum).
		Find(&Videos).Error
	if result != nil {
		return Videos, result
	}
	return Videos, nil
}
