package main

import (
	"douyinOrigin/dao/userDao"
	"douyinOrigin/router"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	//初始化项目依赖
	initDeps()
	//创建一个默认路由
	r := gin.Default()

	//InitRouter(r)
	router.InitRouter(r)

	// 启动HTTP服务，这里端口改为 9090
	err := r.Run(":9090")
	if err != nil {
		log.Panicln("runErr: ", err)
	}
}

// 加载项目依赖
func initDeps() {
	//初始化MySQL数据库
	userDao.InitMySQL()
}
