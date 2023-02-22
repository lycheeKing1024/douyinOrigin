package controller

import (
	"douyinOrigin/config"
	"douyinOrigin/service/userService"
	VideoService "douyinOrigin/service/vedioService"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	Response
	NextTime   int64                `json:"next_time"`
	TableVideo []VideoService.Video `json:"video_list"`
}

// Feed 视频流接口 /feed/
func Feed(c *gin.Context) {
	inputTime := c.Query("latest_time")
	fmt.Println("请求到的参数为：", inputTime)
	var lastTime time.Time
	if inputTime != "0" {
		//parseInt, _ := strconv.ParseInt(inputTime, 10, 64)
		lastTime, _ = time.Parse(config.TimeLayout, inputTime)
		//lastTime = time.Unix(parseInt, 0)
	} else {
		lastTime = time.Now()
	}
	fmt.Println("获取到时间戳为：", lastTime)

	//通过token获得userid
	tokenString := c.Query("token")
	fmt.Println(tokenString)
	myClaims, _ := userService.UserServiceImpl{}.ParseToken(tokenString) //解析token
	userId, err2 := strconv.ParseInt(myClaims.ID, 10, 64)                //通过解析token，拿到userid
	if err2 != nil {
		fmt.Println("解析token 失败，没有拿到userid")
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
			TableVideo: feed,
			NextTime:   nextTime.Unix(),
		})
	}

}

// 投稿接口 /publish/action/
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
	myClaims, _ := userService.UserServiceImpl{}.ParseToken(tokenString) //解析token
	userId, err2 := strconv.ParseInt(myClaims.ID, 10, 64)                //通过解析token，拿到userid
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
