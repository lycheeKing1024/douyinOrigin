package controller

import (
	"douyinOrigin/middleware/jwt"
	"douyinOrigin/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FavoriteListResponse struct {
	Response
	VideoList []service.Video `json:"video_list"` // 用户点赞视频列表
}

// FavoriteAction 点赞/取消点赞操作操作
func FavoriteAction(c *gin.Context) {
	video_id := c.Query("video_id")
	action_type := c.Query("action_type")
	videoId, _ := strconv.ParseInt(video_id, 10, 64)
	actionType, _ := strconv.ParseInt(action_type, 10, 64)

	//通过token获取userId
	tokenString := c.Query("token")
	myClaims, _ := jwt.ParseToken(tokenString)            //解析token
	userId, err2 := strconv.ParseInt(myClaims.ID, 10, 64) //通过解析token，拿到userid
	if err2 != nil {
		fmt.Println("解析token 失败，赞操作中没有拿到userid")
	}
	lsi := service.LikeServiceImpl{}
	//lsi := GetVideo()
	err := lsi.FavouriteAction(userId, videoId, int32(actionType))
	if err != nil {
		fmt.Printf("lsi.FavouriteAction() 失败，%v\n", err)
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "favourite action fail",
		})
	} else {
		c.JSON(http.StatusOK, Response{
			StatusCode: 0,
			StatusMsg:  "favourite action success",
		})
	}
}

// GetFavoriteList 获取点赞列表;
func GetFavoriteList(c *gin.Context) {
	user_id := c.Query("user_id")
	userId, _ := strconv.ParseInt(user_id, 10, 64)
	lsi := service.LikeServiceImpl{}
	//lsi := GetVideo()
	videos, err := lsi.GetFavouriteList(userId, userId)
	if err != nil {
		fmt.Printf("方法like.GetFavouriteList(userid) 失败：%v\n", err)
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "get favouriteList fail",
		})
	} else {
		c.JSON(http.StatusOK, FavoriteListResponse{
			Response:  Response{StatusCode: 0, StatusMsg: "get favouriteList success"},
			VideoList: videos,
		})
	}

}
