package middleware

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

// Ctx Background返回一个非空的Context。 它永远不会被取消，没有值，也没有期限。
// 它通常在main函数，初始化和测试时使用，并用作传入请求的顶级上下文。
var Ctx = context.Background()

var RdbFollowers *redis.Client
var RdbFollowing *redis.Client
var RdbFollowingPart *redis.Client

var RdbLikeUserId *redis.Client  //key:userId,value:VideoId
var RdbLikeVideoId *redis.Client //key:VideoId,value:userId

var RdbVCid *redis.Client //redis db11 -- video_id + comment_id
var RdbCVid *redis.Client //redis db12 -- comment_id + video_id

// InitRedis 初始化Redis连接。
func InitRedis() {
	RdbFollowers = redis.NewClient(&redis.Options{
		Addr:     "1.15.97.114:6379",
		Password: "yangming666@redis",
		DB:       5, // 粉丝列表信息存入 DB5.
	})
	RdbFollowing = redis.NewClient(&redis.Options{
		Addr:     "1.15.97.114:6379",
		Password: "yangming666@redis",
		DB:       6, // 关注列表信息信息存入 DB6.
	})
	RdbFollowingPart = redis.NewClient(&redis.Options{
		Addr:     "1.15.97.114:6379",
		Password: "yangming666@redis",
		DB:       7, // 当前用户是否关注了自己粉丝信息存入 DB7.
	})

	RdbLikeUserId = redis.NewClient(&redis.Options{
		Addr:     "1.15.97.114:6379",
		Password: "yangming666@redis",
		DB:       8, //  选择将点赞视频id信息存入 DB8.
	})

	RdbLikeVideoId = redis.NewClient(&redis.Options{
		Addr:     "1.15.97.114:6379",
		Password: "yangming666@redis",
		DB:       9, //  选择将点赞用户id信息存入 DB9.
	})
	RdbVCid = redis.NewClient(&redis.Options{
		Addr:     "1.15.97.114:6379",
		Password: "yangming666@redis",
		DB:       11, // 选择将video_id中的评论id s存入 DB11.
	})

	RdbCVid = redis.NewClient(&redis.Options{
		Addr:     "1.15.97.114:6379",
		Password: "yangming666@redis",
		DB:       12, // lsy 选择将comment_id对应video_id存入 DB12.
	})
	rdb := redis.NewClient(&redis.Options{
		Addr:     "1.15.97.114:6379",
		Password: "yangming666@redis",
		DB:       0,
	})
	pong, err := rdb.Ping(Ctx).Result()
	if err != nil {
		fmt.Printf("连接redis出错，错误信息：%v", err)
	}
	fmt.Printf("成功连接redis,%v", pong)

}
