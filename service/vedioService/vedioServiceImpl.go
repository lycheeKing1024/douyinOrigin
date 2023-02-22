package VideoService

import (
	"douyinOrigin/config"
	"douyinOrigin/dao/videoDao"
	"douyinOrigin/middleware"
	"douyinOrigin/service/userService"
	"fmt"
	"mime/multipart"
	"sync"
	"time"

	"github.com/google/uuid"
)

type VideoServiceImpl struct {
	userService.UserService
}

// 将数据进行拷贝和转换*（添加作者等信息）
func (vsi VideoServiceImpl) copyVideos(result *[]Video, data *[]videoDao.TableVideo, userId int64) error {
	for _, temp := range *data {
		var video Video
		//将video进行组装，添加一些信息
		vsi.createVideo(&video, &temp, userId)
		*result = append(*result, video)
	}
	return nil
}

func (vsi VideoServiceImpl) createVideo(video *Video, data *videoDao.TableVideo, userId int64) {
	//	建立协程组,当这一组协程全部完成，才会结束本方法
	var wg sync.WaitGroup
	wg.Add(1) //启动4个协程
	var err error
	video.TableVideo = *data

	//插入Author，这里需要将视频的发布者和当前登录的用户传入，才能正确获得isFollow，
	//如果出现错误，不能直接返回失败，将默认值返回，保证稳定
	go func() {
		video.Author, err = vsi.GetUserByCurId(data.AuthorId, userId)
		fmt.Println("在协程中~~")
		if err != nil {
			fmt.Printf("调用 vsi.GetUserByCurId() 方法失败，%v\n", err)
		} else {
			fmt.Printf("调用 vsi.GetUserByCurId() 方法成功\n")
		}
		defer wg.Done() //结束此进程
	}()

	//插入点赞数量,不将nil向上返回
	//go func() {
	//	wg.Done()
	//
	//}()
	//
	//go func() {
	//	defer wg.Done()
	//
	//}()
	//go func() {
	//	defer wg.Done()
	//
	//}()
	wg.Wait()
}

func (vsi VideoServiceImpl) Feed(lastTime time.Time, userId int64) ([]Video, time.Time, error) {
	//先定义好切片的长度
	videos := make([]Video, 0, config.VideoNum)
	data, err := videoDao.GetVideoByLastTime(lastTime)
	if err != nil {
		fmt.Printf("调用videoDao.GetVideoByLastTime() 方法失败，%v", err)
		return nil, time.Time{}, err
	}
	fmt.Printf("调用videoDao.GetVideoByLastTime() 方法成功，\n%v\n", data)
	err = vsi.copyVideos(&videos, &data, userId)
	if err != nil {
		fmt.Printf("err= %v\n", err)
		return nil, time.Time{}, err
	}
	//	返回数据，同时获得视频最早的时间返回
	return videos, data[len(data)-1].PublishTime, nil
}

func (vsi VideoServiceImpl) GetVideo(VideoId int64, userId int64) (Video, error) {
	//TODO implement me
	panic("implement me")
}

// PublishVideo 将传入的视频流保存在文件服务器中，并存储在mysql表中
func (vsi VideoServiceImpl) PublishVideo(data *multipart.FileHeader, userId int64, title string) error {
	//将视频上传到视频服务器里vsftpd ，保存视频链接
	file, err := data.Open() //打开文件
	if err != nil {
		fmt.Printf("data.open()失败，%v", err)
		return err
	}
	fmt.Println("data.open() 成功")
	//	生成一个一个uuid作为视频名字 v4版本
	videoName := uuid.New().String()
	fmt.Printf("视频名称是%v\n", videoName)
	err = videoDao.VideoFTP(file, videoName)
	if err != nil {
		fmt.Printf("videoDao.VideoFTP()失败，%v\n", err)
		return err
	}
	defer file.Close() //关闭文件

	//在服务器上执行ffmpeg 从视频流中获取第一帧截图，并上传图片服务器，保存图片链接
	imageName := videoName

	//向队列中添加消息
	middleware.Ffchan <- middleware.Ffmsg{
		VideoName: videoName,
		ImageName: imageName,
	}
	//组装信息并持久化
	if err = videoDao.Save(videoName, imageName, userId, title); err != nil {
		fmt.Printf("组装信息并持久化 失败~~~,err:%v\n", err)
		return err
	}
	return nil
}

func (vsi VideoServiceImpl) TableVideo(userId int64, curID int64) ([]Video, error) {
	//TODO implement me
	panic("implement me")
}

func (vsi VideoServiceImpl) GetVideoIdList(userId int64) ([]int64, error) {
	//TODO implement me
	panic("implement me")
}
