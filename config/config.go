package config

import "time"

// Secret 密钥
var Secret = "douyinOrigin"

// VideoNum 每次获取视频流的数量 （要求最大30）
const VideoNum = 5

// ftp服务器地址
const ConUftp = "1.15.97.114:21"
const FtpUser = "uftp"
const FtpPassword = "yangming666@uftp"
const HeartbeatTime = 2 * 60

// PlayUrlPrefix 存储视频的链接
const PlayUrlPrefix = "http://1.15.97.114/videos/"

// CoverUrlPrefix 存储视频封面的链接
const CoverUrlPrefix = "http://1.15.97.114/images/"
const TimeLayout = "2006-01-02 15:04:05"

// HostSSH SSH配置
const HostSSH = "1.15.97.114"
const UserSSH = "uftp"
const PasswordSSH = "yangming666@uftp"
const TypeSSH = "password"
const PortSSH = 22
const MaxMsgCount = 100
const SSHHeartbeatTime = 10 * 60

const IsLike = 1      //点赞
const UnLike = 2      //取消点赞
const LikeAction = 1  //点赞的行为
const Atservicets = 3 //操作数据库的最大尝试次数

/* 时间*/
var OneDayOfHours = 60 * 60 * 24
var OneMinute = 60 * 1
var OneMonth = 60 * 60 * 24 * 30
var OneYear = 365 * 60 * 60 * 24
var ExpireTime = time.Hour * 48 // 设置Redis数据热度消散时间。

const DefaultRedisValue = -1 //redis中key对应的预设值，防脏读

// 评论状态
const ValidComment = 1   //评论状态：有效
const InvalidComment = 2 //评论状态：取消（删除）
