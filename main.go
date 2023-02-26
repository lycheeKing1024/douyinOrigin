package main

import (
	"douyinOrigin/dao"
	"douyinOrigin/middleware"
	"douyinOrigin/middleware/rabbitmq"
	"douyinOrigin/router"
	"log"
	"net/http"
	"net/http/pprof"
	_ "net/http/pprof"
	"os"

	"github.com/gin-gonic/gin"
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
	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/debug/pprof/", pprof.Index)
		mux.HandleFunc("/debug/pprof/heap", pprof.Index)
		mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
		mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
		if err := http.ListenAndServe(":9090", nil); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
}

// 加载项目依赖
func initDeps() {
	//初始化MySQL数据库
	dao.InitMySQL()
	//初始化ftp连接
	middleware.InitFTP()
	//	初始话ssh连接
	middleware.InitSSH()

	// 初始化redis-DB0的连接，follow选择的DB0.
	middleware.InitRedis()
	// 初始化rabbitMQ。
	rabbitmq.InitRabbitMQ()

	// 初始化Like的相关消息队列，并开启消费。
	rabbitmq.InitLikeRabbitMQ()
	//初始化Comment的消息队列，并开启消费
	rabbitmq.InitCommentRabbitMQ()
}
