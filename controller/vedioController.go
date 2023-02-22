package controller

import (
	"douyinOrigin/middleware"
	VideoService "douyinOrigin/service/vedioService"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	Response
	NextTime  int64                `json:"next_time"`
	VideoList []VideoService.Video `json:"video_list"`
}
type VideoListResponse struct {
	Response
	VideoList []VideoService.Video `json:"video_list"`
}

// Feed 视频流接口 /feed/
func Feed(c *gin.Context) {
	//inputTime := c.Query("latest_time")
	//fmt.Println("请求到的参数为：", inputTime)
	//var lastTime time.Time
	//if inputTime != "0" {
	//	parseInt, _ := strconv.ParseInt(inputTime, 10, 64)
	//	//lastTime, _ = time.Parse(inputTime, config.TimeLayout)
	//	lastTime = time.Unix(parseInt, 0)
	//} else {
	//	lastTime = time.Now()
	//}
	lastTime := time.Now()
	fmt.Println("获取到时间戳为：", lastTime)

	//通过token获得userid
	tokenString := c.Query("token")
	fmt.Println(tokenString)
	var userId int64
	if tokenString != "" {
		myClaims, _ := middleware.ParseToken(tokenString) //解析token
		userId, _ = strconv.ParseInt(myClaims.ID, 10, 64) //通过解析token，拿到userid
		//if err2 != nil {
		//	fmt.Println("解析token 失败，没有拿到userid")
		//}
	} else {
		userId = 0
	}

	vsi := VideoService.VideoServiceImpl{}
	feed, nextTime, err := vsi.Feed(lastTime, userId)
	if err != nil {
		c.JSON(http.StatusOK, FeedResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "获取视频流失败",
			}})
		return
	} else {
		c.JSON(http.StatusOK, FeedResponse{
			Response: Response{
				StatusCode: 0,
			},
			VideoList: feed,
			NextTime:  nextTime.Unix(),
		})
	}

}

// Publish 投稿接口 /publish/action/
func Publish(c *gin.Context) {
	file, err := c.FormFile("data")
	if err != nil {
		fmt.Printf("获取视频流数据失败：%v", err)
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	//	获取视频标题
	title := c.PostForm("title")
	fmt.Println("title：", title)

	//获取token
	tokenString := c.PostForm("token")
	fmt.Println(tokenString)
	myClaims, _ := middleware.ParseToken(tokenString)     //解析token
	userId, err2 := strconv.ParseInt(myClaims.ID, 10, 64) //通过解析token，拿到userid
	if err2 != nil {
		fmt.Println("解析token 失败，没有拿到userid")
	}

	vsi := VideoService.VideoServiceImpl{}

	err3 := vsi.PublishVideo(file, userId, title)
	if err3 != nil {
		//log.Printf("vsi.Publish() 失败：%v\n", err)
		fmt.Printf("vsi.Publish() 失败：%v\n", err3)
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	//log.Printf("vsi.Publish() 成功")
	fmt.Println("vsi.Publish() 成功")

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  "uploaded successfully",
	})

}

// PublishList 投稿列表 /publish/list
func PublishList(c *gin.Context) {
	user_id := c.Query("user_id")
	userId, _ := strconv.ParseInt(user_id, 10, 64)
	//log.Panicf("当前用户id= %v\n", userId)
	vsi := VideoService.VideoServiceImpl{}
	list, err := vsi.List(userId)
	if err != nil {
		log.Printf("调用videoService.List(%v)出现错误：%v\n", userId, err)
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{StatusCode: 1, StatusMsg: "获取视频列表失败"},
		})
		return
	}
	c.JSON(http.StatusOK, VideoListResponse{
		Response:  Response{StatusCode: 0, StatusMsg: "获取视频列表成功"},
		VideoList: list,
	})
}
