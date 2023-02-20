package controller

import (
	VideoService "douyinOrigin/service/vedioService"
	"fmt"
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

// Feed 视频流接口 /douyin/feed/
func Feed(c *gin.Context) {
	inputTime := c.Query("latest_time")
	fmt.Println("请求到的参数为：", inputTime)
	var lastTime time.Time
	if inputTime != "0" {
		parseInt, _ := strconv.ParseInt(inputTime, 10, 64)
		lastTime = time.Unix(parseInt, 0)
	} else {
		lastTime = time.Now()
	}
	fmt.Println("获取到时间戳为：", lastTime)
	userId, _ := strconv.ParseInt(c.GetString("userId"), 10, 64)
	fmt.Println("获取到的用户id：", userId)
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
