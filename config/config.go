package config

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
const PlayUrlPrefix = "http://1.15.97.114/videos"

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
