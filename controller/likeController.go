package controller

import (
	"douyinOrigin/middleware"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// FavoriteAction 点赞/取消点赞操作操作
func FavoriteAction(c *gin.Context) {
	video_id := c.Query("video_id")
	action_type := c.Query("action_type")
	videoId, _ := strconv.ParseInt(video_id, 10, 64)
	actionType, _ := strconv.ParseInt(action_type, 10, 64)
	//通过token获取userId
	tokenString := c.Query("token")
	myClaims, _ := middleware.ParseToken(tokenString)     //解析token
	userId, err2 := strconv.ParseInt(myClaims.ID, 10, 64) //通过解析token，拿到userid
	if err2 != nil {
		fmt.Println("解析token 失败，赞操作中没有拿到userid")
	}

}
