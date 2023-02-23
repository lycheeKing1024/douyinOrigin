package dao

import (
	"douyinOrigin/config"
	"douyinOrigin/middleware"
	"fmt"
	"io"
	"time"
)

type TableVideo struct {
	Id          int64  `json:"id"` //视频id
	AuthorId    int64  //视频作者信息
	PlayUrl     string `json:"play_url"`  //视频播放地址
	CoverUrl    string `json:"cover_url"` //视频封面地址
	PublishTime time.Time
	Title       string `json:"title"` //视频标题
}

// TableName 将 TableVideo 结构体映射到数据表 videos
func (TableVideo) TableName() string {
	return "videos"
}

// GetVideoByVeidoId 根据视频id获取视频信息
func GetVideoByVideoId(VideoId int64) (TableVideo, error) {
	var TableVideo TableVideo
	TableVideo.Id = VideoId
	result := SqlSession.First(&TableVideo)
	if result.Error != nil {
		return TableVideo, result.Error
	}
	return TableVideo, nil
}

// GetVideosByAuthorId 根据作者id查询视频信息，返回切片
func GetVideosByAuthorId(authorId int64) ([]TableVideo, error) {
	var TableVideos []TableVideo
	result := SqlSession.Where("author_id=?", authorId).Find(&TableVideos)
	if result.Error != nil {
		//如果出现问题，返回对应到空，并且返回错误信息
		return nil, result.Error
	}
	return TableVideos, nil
}

// GetVideoByLastTime 根据传入的时间来获取这个时间之前的一些数据
func GetVideoByLastTime(lastTime time.Time) ([]TableVideo, error) {
	Videos := make([]TableVideo, 0, config.VideoNum)
	result := SqlSession.Where("publish_time<=?", lastTime).Order("publish_time desc").Limit(config.VideoNum).
		Find(&Videos).Error
	if result != nil {
		return Videos, result
	}
	return Videos, nil
}

// VideoFTP 通过ftp将视频传入服务器
func VideoFTP(file io.Reader, videoName string) error {
	path, _ := middleware.MyFTP.Pwd()
	fmt.Printf("path= %v", path)
	//	cwd 转到视频目录下
	if err := middleware.MyFTP.Cwd("./videos"); err != nil {
		fmt.Printf("进入 videos 目录失败,err：%v\n", err)
		return err
	}
	if err2 := middleware.MyFTP.Stor(videoName+".mp4", file); err2 != nil {
		fmt.Printf("上传 视频 失败！！！,err2：%v\n", err2)
		return err2
	}
	fmt.Println("上传 视频 成功！！！")
	return nil

}

// ImageFTP 通过ftp将图片传入服务器
// 将图片传入FTP服务器中，但是这里要注意图片的格式随着名字一起给,同时调用时需要自己结束流
func ImageFTP(file io.Reader, imageName string) error {
	//	cwd 转到视频目录下
	if err := middleware.MyFTP.Cwd("./images"); err != nil {
		fmt.Println("进入 images 目录失败")
		return err
	}
	if err2 := middleware.MyFTP.Stor(imageName, file); err2 != nil {
		fmt.Println("上传 图片 失败！！！")
		return err2
	}
	fmt.Println("上传 图片 成功！！！")
	return nil
}

func Save(videoName string, imageName string, authorId int64, title string) error {
	video := TableVideo{
		AuthorId:    authorId,
		PlayUrl:     config.PlayUrlPrefix + videoName + ".mp4",
		CoverUrl:    config.CoverUrlPrefix + imageName + ".jpg",
		PublishTime: time.Now(),
		Title:       title,
	}
	result := SqlSession.Save(&video)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetVideoIdByAuthorId // 通过作者id来查询发布的视频id切片集合
func GetVideoIdByAuthorId(authorId int64) ([]int64, error) {
	var videoId []int64
	if err := SqlSession.Model(&TableVideo{}).Where("author_id").Pluck("id", &videoId).Error; err != nil {
		fmt.Printf("发布列表获取视频id失败,err：%v\n", err)
		return nil, err
	}
	return videoId, nil
}
