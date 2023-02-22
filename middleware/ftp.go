package middleware

import (
	"douyinOrigin/config"
	"fmt"

	"github.com/dutchcoders/goftp"
)

var MyFTP *goftp.FTP

func InitFTP() {
	var err error
	//    获取ftp连接
	MyFTP, err = goftp.Connect(config.ConUftp)
	if err != nil {
		fmt.Println("Error connecting to FTP server:", err)
		return
	}
	//defer MyFTP.Close()
	fmt.Println("FTP Successfully connected !!")
	if err = MyFTP.Login(config.FtpUser, config.FtpPassword); err != nil {
		//panic(err)
		fmt.Println("FTP 登录失败")
		return
	}
	fmt.Println("FTP 登录成功")
	//	维持长连接
	go keepAlive()
}

//func keepAlive() {
//	time.Sleep(time.Duration(config.HeartbeatTime) * time.Second)
//	MyFTP.Noop()
//}
