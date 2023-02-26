package dao

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	//"gorm.io/gorm/logger"
	//"log"
	//"os"
	//"time"
)

var SqlSession *gorm.DB

func InitMySQL() {
	//newLogger := logger.New(
	//	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	//	logger.Config{
	//		SlowThreshold: time.Second,   // 慢 SQL 阈值
	//		LogLevel:      logger.Silent, // Log level
	//		Colorful:      true,          // 彩色打印
	//	},
	//)

	// 配置MySQL连接参数
	username := "root" //账号
	password := "yangming666@mysql"
	host := "1.15.97.114"
	port := 3306
	dbName := "douyin"

	var err error

	//通过前面的数据库参数，拼接MYSQL DSN， 其实就是数据库连接串（数据源名称）
	//MYSQL dsn格式： {username}:{password}@tcp({host}:{port})/{SqlSessionname}?charset=utf8&parseTime=True&loc=Local
	//类似{username}使用花括号包着的名字都是需要替换的参数
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbName)
	//要支持完整的UTF-8编码，需要将 charset=utf8 改为 charset=utf8mb4
	//	连接mysql
	SqlSession, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		//	这里就使用log打印了
		//log.Panicle("数据库连接失败，error: " + err.Error())
		fmt.Println(err.Error())
	} else {
		//log.Panicle("数据库连接成功")
		fmt.Println("数据库连接成功")
	}
	// 3.程序退出关闭数据库
	// 在v2中,gorm维护了一个连接池，初始化db之后所有的连接都由库来管理。所以不需要使用者手动关闭
	// defer db.Close()
}
