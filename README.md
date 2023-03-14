# 青训营后端结业项目答辩汇报文档



# 一、项目介绍

> 基于Gin+Gorm+MySQL+Redis+RabbitMQ 的极简版抖音服务端
>
> 项目服务地址 http://1.15.97.114:9090/
>
> Github地址 https://github.com/1Carryon/douyinOrigin

# 二、项目分工

| **团队成员** | **主要贡献** |
| ------------ | ------------ |
| 李春旺       | 全部         |

# 三、项目实现

### 3.1 技术选型与相关开发文档

- 选择Go作为项目开发语言

- 数据层选用MySQL作为持久化数据库，redis作为缓存的机制。MySQL到redis有良好的fdw支持。

- 采用Nginx用于负载均衡

- Gin作为本项目中的Web框架，简单易用

- 数据交互采用go的Gorm框架与redis模组，gorm的查询封装已经可以帮助我们抵御sql注入，只要不使用原生sql执行即可。根据gorm的建议，我们不使用连接查询。将每个类型的实体单独拆分为表。同时redis可以将一些信息存进缓存中，当点赞或者取消赞的时候，只需要添加或者删除对应的信息即可，响应时间更快，再采用往rabbitmq中传递需要修改数据库的关键信息，通过协程的方式修改数据库。

- 此外对于权限管理，本项目选择采用jwt进行用户鉴权,所以无需在服务器上存储用户状态。同时自身还可以携带额外信息，如数字签名、时间期限，以增加其安全性。

#### 相关开发文档

gorm[https://gorm.io/zh_CN/docs]

gin[https://gin-gonic.com/zh-cn/docs/]

go-redis[https://redis.uptrace.dev/zh/]

go-jwt[https://pkg.go.dev/github.com/golang-jwt/jwt/v5]

redis[https://redis.io/docs/]

rabbirmq[https://developer.aliyun.com/article/769883]

### 3.2 架构设计

> 可以补充场景分析环节，明确要解决的问题和前提假设，比如预计0.5%的用户属于大V，粉丝很多，也会经常上传视频，当前架构的解决方案是xxx。

本项目采用分层架构，客户端发送请求后，首先进行用户鉴权操作，通过鉴权后，Controller控制层对请求进行响应并解析请求中的数据，将数据传入Service服务层进行数据处理，服务层调用Dao层进行判断用户信息、增加用户。

#### 整体架构

<img src="pictures/整体架构图.jpg" alt="img" style="zoom:50%;" />

#### 数据库的设计

![流程图](pictures/流程图.jpg)

#### Redis架构的设计

<img src="pictures/redis.png" alt="img" style="zoom:50%;" />

#### RabbitMQ架构设计

<img src="pictures/rabbitmq.jpg" alt="img" style="zoom:67%;" />

#### 服务模块的设计

<img src="pictures/service.png" alt="img" style="zoom:67%;" />

### 3.3 项目代码介绍

#### 项目目录结构

```Bash
.
├── config
│   ├── config.go
│   └── douyin.sql
├── controller
│   ├── commentController.go
│   ├── likeController.go
│   ├── userController.go
│   └── videoController.go
├── dao
│   ├── commentDao.go
│   ├── initMySQL.go
│   ├── likeDao.go
│   ├── userDao.go
│   └── videoDao.go
├── .gitignore
├── go.mod
├── goplantuml.exe
├── go.sum
├── LICENSE
├── main.go
├── middleware
│   ├── ffmpeg.go
│   ├── ftp.go
│   ├── jwt
│   │   └── jwt.go
│   ├── rabbitmq
│   │   ├── commentMQ.go
│   │   ├── likeMQ.go
│   │   └── rabbitMQ.go
│   └── redis.go
├── router
│   └── router.go
└── service
    ├── commentService.go
    ├── commentServiceImpl.go
    ├── commentServiceImpl_test.go
    ├── likeService.go
    ├── likeServiceImpl.go
    ├── likeServiceImpl_test.go
    ├── service.puml
    ├── userService.go
    ├── userServiceImpl.go
    ├── userServiceImpl_test.go
    ├── videoService.go
    ├── videoServiceImpl.go
    └── videoServiceImpl_test.go
```

#### 数据库表结构

```Go
type TableUser struct {
   Id       int64 //主键
   Name     string
   Password string
   //NickName string
   Avatar          string // 用户头像
   BackgroundImage string // 用户个人页顶部大图
   Signature       string // 个人简介
   WorkCount       int64  // 作品数
}
type TableVideo struct {
   Id          int64  `json:"id"` //视频id
   AuthorId    int64  //视频作者信息
   PlayUrl     string `json:"play_url"`  //视频播放地址
   CoverUrl    string `json:"cover_url"` //视频封面地址
   PublishTime time.Time
   Title       string `json:"title"` //视频标题
}
type Like struct {
   Id      int64 //自增主键
   UserId  int64 //点赞用户id
   VideoId int64 //视频id
   Cancel  int8  //是否点赞，1为点赞，2为取消赞
}
// 评论信息-数据库中的结构体-dao层使用
type Comment struct {
   Id          int64     //评论id
   UserId      int64     //评论用户id
   VideoId     int64     //视频id
   CommentText string    //评论内容
   CreateDate  time.Time //评论发布的日期mm-dd
   Cancel      int32     //取消评论为1，发布评论为0
}
```

#### 路由结构

```Go
func InitRouter(r *gin.Engine) {
   apiServer := r.Group("/douyin")
   {
      //基础接口
      apiServer.POST("/user/register/", controller.Register)
      apiServer.POST("/user/login/", controller.Login)
      apiServer.GET("/user/", controller.UserInfo)
      apiServer.GET("/feed/", controller.Feed)
      apiServer.POST("/publish/action/", controller.Publish)
      apiServer.GET("/publish/list/", controller.PublishList)

      //互动接口
      apiServer.POST("/favorite/action/", controller.FavoriteAction)
      apiServer.GET("/favorite/list/", controller.GetFavoriteList)
      apiServer.POST("/comment/action/", controller.CommentAction)
      apiServer.GET("/comment/list/", controller.CommentList)

   }
}
```

#### 中间件

##### Jwt

```Go
// EnCoder 密码加密
func EnCoder(password string) string {
   hash := hmac.New(sha256.New, []byte(password)) //创建对应的sha256哈希加密算法
   sha := hex.EncodeToString(hash.Sum(nil))
   fmt.Println("EnCoder: ", sha)
   return sha
}

// GenerateToken 根据username生成一个token
func GenerateToken(username string) string {
   //u := UserService.GetTableUserByUsername(new(UserServiceImpl), username)
   u, _ := dao.GetTableUserByUsername(username)
   fmt.Printf("generatetoken: %v\n", u)
   token := NewToken(u)
   println(token)
   return token
}

// NewToken 根据信息创建token
func NewToken(u dao.TableUser) string {
   expiresTime := time.Now().Add(time.Hour * time.Duration(12)) //设置过期时间为12小时
   fmt.Printf("expiresTime: %v\n", expiresTime)
   id64 := u.Id
   fmt.Printf("id: %v\n", strconv.FormatInt(id64, 10))
   claims := MyClaims{
      jwt.RegisteredClaims{
         Audience:  jwt.ClaimStrings{u.Name},        //受众
         ExpiresAt: jwt.NewNumericDate(expiresTime), //过期时间
         ID:        strconv.FormatInt(id64, 10),     //编号
         IssuedAt:  jwt.NewNumericDate(time.Now()),  //签发时间
         Issuer:    "yangming",                      //签发人
         NotBefore: jwt.NewNumericDate(time.Now()),  //生效时间
         Subject:   "token",                         //主题
      },
   }

   tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
   if token, err := tokenClaims.SignedString(jwtSecret); err == nil {
      token = "Bearer " + token
      println("generate token success!\n")
      return token
   } else {
      println("generate token fail\n")
      return "fail"
   }
}

type MyClaims struct {
   jwt.RegisteredClaims // 注意!这是jwt-go的v4版本新增的，原先是jwt.StandardClaims
}

var jwtSecret = []byte(config.Secret) //jwt密钥
// ParseToken 解析token
func ParseToken(tokenString string) (*MyClaims, error) {
   //tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
   prefixLen := len("Bearer ")
   tokenString = tokenString[prefixLen:]
   //fmt.Println(tokenString)
   tokenClaims, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
      return jwtSecret, nil
   })
   if err != nil {
      fmt.Printf("解析token失败！！！,err：%v", err)
   }
   if claims, ok := tokenClaims.Claims.(*MyClaims); ok && tokenClaims.Valid {
      return claims, nil
   }
   return nil, err
}
```

##### Ffmpeg

```Go
// InitSSH 建立SSH客户端，但是会不会超时导致无法链接，这个需要做一些措施
func InitSSH() {
   var err error
   //创建sshp登陆配置
   SSHconfig := &ssh.ClientConfig{
      Timeout:         5 * time.Second, //ssh 连接time out 时间一秒钟, 如果ssh验证错误 会在一秒内返回
      User:            config.UserSSH,
      HostKeyCallback: ssh.InsecureIgnoreHostKey(), //这个可以, 但是不够安全

      //HostKeyCallback: hostKeyCallBackFunc(h.Host),
   }
   if config.TypeSSH == "password" {
      SSHconfig.Auth = []ssh.AuthMethod{ssh.Password(config.PasswordSSH)}
   }
   //dial 获取ssh client
   addr := fmt.Sprintf("%s:%d", config.HostSSH, config.PortSSH)
   ClientSSH, err = ssh.Dial("tcp", addr, SSHconfig)
   if err != nil {
      log.Fatal("创建ssh client 失败", err)
   }
   log.Printf("获取到客户端：%v", ClientSSH)
   //建立通道，作为队列使用,并且确立缓冲区大小
   Ffchan = make(chan Ffmsg, config.MaxMsgCount)
   //建立携程用于派遣
   go dispatcher()
   go keepAlive()
}

// 通过增加携程，将获取的信息进行派遣，当信息处理失败之后，还会将处理方式放入通道形成的队列中
func dispatcher() {
   for ffmsg := range Ffchan {
      go func(f Ffmsg) {
         err := Ffmpeg(f.VideoName, f.ImageName)
         if err != nil {
            Ffchan <- f
            log.Fatal("派遣失败：重新派遣")
         }
         log.Printf("视频%v截图处理成功", f.VideoName)
      }(ffmsg)
   }
}

// Ffmpeg 通过远程调用ffmpeg命令来创建视频截图
func Ffmpeg(videoName string, imageName string) error {
   session, err := ClientSSH.NewSession()
   if err != nil {
      log.Fatal("创建ssh session 失败", err)
   }
   defer session.Close()
   //执行远程命令 ffmpeg -ss 00:00:01 -i /home/uftp/videos/1.mp4 -vframes 1 /home/uftp/images/4.jpg
   combo, err := session.CombinedOutput("ls;/usr/local/ffmpeg/bin/ffmpeg -ss 00:00:01 -i /home/uftp/videos/" + videoName + ".mp4 -vframes 1 /home/uftp/images/" + imageName + ".jpg")
   if err != nil {
      //log.Fatal("远程执行cmd 失败", err)
      log.Fatal("命令输出:", string(combo))
      return err
   }
   //fmt.Println("命令输出:", string(combo))
   return nil
}

// 维持长链接
func keepAlive() {
   time.Sleep(time.Duration(config.SSHHeartbeatTime) * time.Second)
   session, _ := ClientSSH.NewSession()
   session.Close()
}
```

##### Ftp

```Go
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
   // 维持长连接
   go keepAlive()
}
```

##### Redis

```Go
/ InitRedis 初始化Redis连接。
func InitRedis() {
   RdbFollowers = redis.NewClient(&redis.Options{
      DB:       5, // 粉丝列表信息存入 DB5.
   })
   RdbFollowing = redis.NewClient(&redis.Options{
      DB:       6, // 关注列表信息信息存入 DB6.
   })
   RdbFollowingPart = redis.NewClient(&redis.Options{
      DB:       7, // 当前用户是否关注了自己粉丝信息存入 DB7.
   })

   RdbLikeUserId = redis.NewClient(&redis.Options{
      DB:       8, //  选择将点赞视频id信息存入 DB8.
   })

   RdbLikeVideoId = redis.NewClient(&redis.Options{
      DB:       9, //  选择将点赞用户id信息存入 DB9.
   })
   RdbVCid = redis.NewClient(&redis.Options{
      DB:       11, // 选择将video_id中的评论id s存入 DB11.
   })

   RdbCVid = redis.NewClient(&redis.Options{
      Addr:     "",
      Password: "",
      DB:       12, // lsy 选择将comment_id对应video_id存入 DB12.
   })
   rdb := redis.NewClient(&redis.Options{
      Addr:     "",
      Password: "",
      DB:       0,
   })
   pong, err := rdb.Ping(Ctx).Result()
   if err != nil {
      fmt.Printf("连接redis出错，错误信息：%v", err)
   }
   fmt.Printf("成功连接redis,%v", pong)

}
```

##### Rabbitmq

部分连接层次代码

```Go
// InitRabbitMQ 初始化RabbitMQ的连接和通道。
func InitRabbitMQ() {

   Rmq = &RabbitMQ{
      mqurl: MQURL,
   }
   dial, err := amqp.Dial(Rmq.mqurl)
   Rmq.failOnErr(err, "创建连接失败")
   Rmq.conn = dial

}

// 连接出错时，输出错误信息。
func (r *RabbitMQ) failOnErr(err error, message string) {
   if err != nil {
      fmt.Printf("%s:%s\n", err, message)
      //panic(fmt.Sprintf("%s:%s\n", err, message))
   }
}

// 关闭mq通道和mq的连接。
func (r *RabbitMQ) destroy() {
   r.conn.Close()
}
```

#### 服务接口

##### 用户模块

```Go
type UserService interface {
   // GetTableUserList 获取TableUser对象
   GetTableUserList() []dao.TableUser

   // GetTableUserByUsername 根据user_name获得User对象
   GetTableUserByUsername(name string) dao.TableUser

   // GetTableUserById 根据 user_id 获得User对象
   GetTableUserById(id int64) dao.TableUser

   // InsertTableUser 将user插入到数据表中
   InsertTableUser(tableUser *dao.TableUser) bool

   // GetUserByCurId 在登录情况（curId）下根据user_id获取User对象，curId判断是否点赞
   GetUserByCurId(id int64, curId int64) (User, error)
}
```

##### 视频模块

```Go
type VideoService interface {
   // Feed 通过传入时间戳，当前用户的id，返回对应的视频切片，以及视频最早的发布时间
   Feed(lastTime time.Time, userId int64) ([]Video, time.Time, error)

   // 传入视频id获得对应的视频对象
   GetVideo(videoId int64, userId int64) (Video, error)

   // PublishVideo 将传入的视频流保存到文件服务器中，并存储在mysql中
   PublishVideo(data *multipart.FileHeader, userId int64, title string) error

   // List TableVideo 通过userId查询对应用户发布的视频，并返回视频数组
   List(userId int64) ([]Video, error)

   // GetVideoIdList 通过一个作者id，返回用户发布的视频id切片数组
   GetVideoIdList(userId int64) ([]int64, error)
}
```

##### 点赞模块

```Go
type LikeService interface {

   // IsFavourite 根据当前视频id判断是否点赞了该视频。
   IsFavourite(videoId int64, userId int64) (bool, error)

   // FavouriteCount 根据当前视频id获取当前视频点赞数量。
   FavouriteCount(videoId int64) (int64, error)

   // TotalFavourite 根据userId获取这个用户总共被点赞数量
   TotalFavourite(userId int64) (int64, error)

   // FavouriteVideoCount 根据userId获取这个用户点赞视频数量
   FavouriteVideoCount(userId int64) (int64, error)

   // FavouriteAction 当前用户对视频的点赞操作 ,并把这个行为更新到like表中。
   //当前操作行为，1点赞，2取消点赞。
   FavouriteAction(userId int64, videoId int64, actionType int32) error

   // GetFavouriteList 获取当前用户的所有点赞视频，调用videoService的方法
   GetFavouriteList(userId int64, curId int64) ([]Video, error)
}
```

##### 评论模块

```Go
// CommentService 接口定义
// 发表评论-使用的结构体-service层引用dao层↑的Comment。
type CommentService interface {

   // CountFromVideoId 1.根据videoId获取视频评论数量的接口
   CountFromVideoId(id int64) (int64, error)

   // Send 2、发表评论，传进来评论的基本信息，返回保存是否成功的状态描述
   Send(comment dao.Comment) (CommentInfo, error)

   // DelComment 3、删除评论，传入评论id即可，返回错误状态信息
   DelComment(commentId int64) error

   // GetList 4、查看评论列表-返回评论list-在controller层再封装外层的状态信息
   GetList(videoId int64, userId int64) ([]CommentInfo, error)
}
```

# 四、测试结果

> 建议从功能测试和性能测试两部分分析，其中功能测试补充测试用例，性能测试补充性能分析报告、可优化点等内容。

### **功能测试**

#### 评论模块

```Bash
=== RUN   TestIsFavourite
数据库连接成功
FTP Successfully connected !!
FTP 登录成功
2023/02/26 14:16:28 获取到客户端：&{0xc0001d7580 {0 {0 0}}        {{0 0} []} {0 0} map[]}
成功连接redis,PONG2023/02/26 14:16:29 [*] Waiting for messagees    ,To exit press CTRL+C
2023/02/26 14:16:29 [*] Waiting for messagees,To exit press CTRL+C
方法:IsFavourite RedisLikeUserId query value成功false--- PASS:     TestIsFavourite (1.91s)
PASS
=== RUN   TestFavouriteCount
数据库连接成功
FTP Successfully connected !!
FTP 登录成功
2023/02/26 14:18:28 获取到客户端：&{0xc0001d7580 {0 {0        0}} {{0 0} []} {0 0} map[]}
成功连接redis,PONG2023/02/26 14:18:28 [*] Waiting for messag    ees,To exit press CTRL+C
2023/02/26 14:18:28 [*] Waiting for messagees,To exit press CTRL+C
方法:FavouriteCount RedisLikeVideoId query count 成功0--- PA    SS: TestFavouriteCount (1.94s)
PASS
=== RUN   TestTotalFavourite
数据库连接成功
FTP Successfully connected !!
FTP 登录成功
2023/02/26 14:18:43 获取到客户端：&{0xc0001d7580 {0 {0        0}} {{0 0} []} {0 0} map[]}
成功连接redis,PONGuserid=3
2023/02/26 14:18:44 [*] Waiting for messagees,To exit press CTRL+C
2023/02/26 14:18:44 [*] Waiting for messagees,To exit press CTRL+C
方法:FavouriteCount RedisLikeVideoId query count 成功方法:Fa      vouriteCount RedisLikeVideoId query count 成功方法:FavouriteCount RedisLikeVideo    Id query count 成功方法:FavouriteCount RedisLikeVideoId query count 成功方法:Fav        ouriteCount RedisLikeVideoId query count 成功方法:FavouriteCount RedisLikeVideoI    d query count 成功方法:FavouriteCount RedisLikeVideoId query count 成功方法:Favo        uriteCount RedisLikeVideoId query count 成功方法:FavouriteCount RedisLikeVideoId     query count 成功方法:FavouriteCount RedisLikeVideoId query count 成功方法:Favou        riteCount RedisLikeVideoId query count 成功方法:FavouriteCount RedisLikeVideoId     query count 成功10--- PASS: TestTotalFavourite (1.86s)
PASS
```

#### 用户模块

```Bash
Running tool: /usr/local/go/bin/go test -timeout 30s -run ^(TestFeed|TestUserAction|TestPublish)$ douyinOrigin/test -v -count=1

=== RUN   TestFeed
    /home/go/douyinOrigin/test/printer.go:116: GET /douyin/feed/ HTTP/1.1
        Host: localhost:8080

    /home/go/douyinOrigin/test/printer.go:133: HTTP/1.1 200 OK 8.972044ms
        Content-Length: 384
        Content-Type: application/json; charset=utf-8
        Date: Tue, 21 Feb 2023 13:40:58 GMT

        {"status_code":0,"video_list":[{"id":1,"author":{"id":1,"name":"test","work_count":1},"play_url":"http://47.113.179.123:8080/static/1_003198fdd79363536c7c6db516d174df06ed9993879b7eb2234269c51444f833.mp4","cover_url":"http://47.113.179.123:8080/static/1_003198fdd79363536c7c6db516d174df06ed9993879b7eb2234269c51444f833.jpg","is_favorite":false,"title":"heusb"}],"next_time":1676986858}
--- PASS: TestFeed (0.01s)
=== RUN   TestUserAction
    /home/go/douyinOrigin/test/printer.go:116: POST /douyin/user/register/?password=douyin47354&username=douyin47354 HTTP/1.1
        Host: localhost:8080
        Content-Type: application/x-www-form-urlencoded

        password=douyin47354&username=douyin47354
    /home/go/douyinOrigin/test/printer.go:133: HTTP/1.1 200 OK 2.699258ms
        Content-Length: 196
        Content-Type: application/json; charset=utf-8
        Date: Tue, 21 Feb 2023 13:40:58 GMT

        {"status_code":0,"user_id":2,"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjoyLCJ1c2VyIjoyLCJleHAiOjE2NzcwNzMyNTgsImlhdCI6MTY3Njk4Njg1OH0.3zqSoKgIBRDUNcXl8YaSjgmGjdB8CS8pjnIOb_3HP4I"}
    /home/go/douyinOrigin/test/printer.go:116: POST /douyin/user/login/?password=douyin47354&username=douyin47354 HTTP/1.1
        Host: localhost:8080
        Content-Type: application/x-www-form-urlencoded

        password=douyin47354&username=douyin47354
    /home/go/douyinOrigin/test/printer.go:133: HTTP/1.1 200 OK 1.405326ms
        Content-Length: 196
        Content-Type: application/json; charset=utf-8
        Date: Tue, 21 Feb 2023 13:40:58 GMT

        {"status_code":0,"user_id":2,"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjozLCJ1c2VyIjoyLCJleHAiOjE2NzcwNzMyNTgsImlhdCI6MTY3Njk4Njg1OH0.Px3odJDUuezxYg7cUlJrCX7W25-2m6PrnU01OXQQw2s"}
    /home/go/douyinOrigin/test/printer.go:116: GET /douyin/user/?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjozLCJ1c2VyIjoyLCJleHAiOjE2NzcwNzMyNTgsImlhdCI6MTY3Njk4Njg1OH0.Px3odJDUuezxYg7cUlJrCX7W25-2m6PrnU01OXQQw2s HTTP/1.1
        Host: localhost:8080

    /home/go/douyinOrigin/test/printer.go:133: HTTP/1.1 200 OK 3.540688ms
        Content-Length: 54
        Content-Type: application/json; charset=utf-8
        Date: Tue, 21 Feb 2023 13:40:58 GMT

        {"status_code":0,"user":{"id":2,"name":"douyin47354"}}
--- PASS: TestUserAction (0.01s)
PASS
ok      douyinOrigin/test  0.372s
```

#### 视频模块

```Bash
=== RUN   TestPublish
    /home/go/douyinOrigin/test/printer.go:116: POST /douyin/user/register/?password=douyinTestUserA&username=douyinTestUserA HTTP/1.1
        Host: localhost:8080
        Content-Type: application/x-www-form-urlencoded

        password=douyinTestUserA&username=douyinTestUserA
    /home/go/douyinOrigin/test/printer.go:133: HTTP/1.1 200 OK 3.646741ms
        Content-Length: 196
        Content-Type: application/json; charset=utf-8
        Date: Tue, 21 Feb 2023 13:40:58 GMT

        {"status_code":0,"user_id":3,"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjo0LCJ1c2VyIjozLCJleHAiOjE2NzcwNzMyNTgsImlhdCI6MTY3Njk4Njg1OH0.BjPjXohAGE1fRHTG_EDpXH-jyMMqVy-tERIAzXXV3OQ"}
    /home/go/douyinOrigin/test/printer.go:116: POST /douyin/publish/action/ HTTP/1.1
        Host: localhost:8080
        Content-Type: multipart/form-data; boundary=6f3a1ef51cb3e1e0989a5e7bf0050155a90c91577a7c20ce46d0415293f9

        --6f3a1ef51cb3e1e0989a5e7bf0050155a90c91577a7c20ce46d0415293f9
        Content-Disposition: form-data; name="data"; filename="../public/bear.mp4"
        Content-Type: application/octet-stream

        [file data content, skip]

        Content-Disposition: form-data; name="token"

        eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjo0LCJ1c2VyIjozLCJleHAiOjE2NzcwNzMyNTgsImlhdCI6MTY3Njk4Njg1OH0.BjPjXohAGE1fRHTG_EDpXH-jyMMqVy-tERIAzXXV3OQ
        --6f3a1ef51cb3e1e0989a5e7bf0050155a90c91577a7c20ce46d0415293f9
        Content-Disposition: form-data; name="title"

        Bear
        --6f3a1ef51cb3e1e0989a5e7bf0050155a90c91577a7c20ce46d0415293f9--
    /home/go/douyinOrigin/test/printer.go:133: HTTP/1.1 200 OK 220.060574ms
        Content-Length: 125
        Content-Type: application/json; charset=utf-8
        Date: Tue, 21 Feb 2023 13:40:59 GMT

        {"status_code":0,"status_msg":"3_3472a10a299cfd24a062a2f9415e57e3fd3fad14df59e9182571fef592e4eb5f.mp4 uploaded successfully"}
    /home/go/douyinOrigin/test/printer.go:116: GET /douyin/publish/list/?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjo0LCJ1c2VyIjozLCJleHAiOjE2NzcwNzMyNTgsImlhdCI6MTY3Njk4Njg1OH0.BjPjXohAGE1fRHTG_EDpXH-jyMMqVy-tERIAzXXV3OQ&user_id=3 HTTP/1.1
        Host: localhost:8080

    /home/go/douyinOrigin/test/printer.go:133: HTTP/1.1 200 OK 2.656585ms
        Content-Length: 371
        Content-Type: application/json; charset=utf-8
        Date: Tue, 21 Feb 2023 13:40:59 GMT

        {"status_code":0,"video_list":[{"id":2,"author":{"id":3,"name":"douyinTestUserA","work_count":1},"play_url":"http://47.113.179.123:8080/static/3_3472a10a299cfd24a062a2f9415e57e3fd3fad14df59e9182571fef592e4eb5f.mp4","cover_url":"http://47.113.179.123:8080/static/3_3472a10a299cfd24a062a2f9415e57e3fd3fad14df59e9182571fef592e4eb5f.jpg","is_favorite":false,"title":"Bear"}]}
--- PASS: TestPublish (0.35s)
```

#### 点赞模块

```Bash
Running tool: /usr/local/go/bin/go test -timeout 30s -run ^(TestFavorite|TestComment)$ douyinOrigin/test -v -count=1

=== RUN   TestFavorite
    /home/go/douyinOrigin/test/printer.go:116: GET /douyin/feed/ HTTP/1.1
        Host: localhost:8080

    /home/go/douyinOrigin/test/printer.go:133: HTTP/1.1 200 OK 5.531507ms
        Content-Length: 723
        Content-Type: application/json; charset=utf-8
        Date: Tue, 21 Feb 2023 13:44:48 GMT

        {"status_code":0,"video_list":[{"id":2,"author":{"id":3,"name":"douyinTestUserA","work_count":1},"play_url":"http://47.113.179.123:8080/static/3_3472a10a299cfd24a062a2f9415e57e3fd3fad14df59e9182571fef592e4eb5f.mp4","cover_url":"http://47.113.179.123:8080/static/3_3472a10a299cfd24a062a2f9415e57e3fd3fad14df59e9182571fef592e4eb5f.jpg","is_favorite":false,"title":"Bear"},{"id":1,"author":{"id":1,"name":"test","work_count":1},"play_url":"http://47.113.179.123:8080/static/1_003198fdd79363536c7c6db516d174df06ed9993879b7eb2234269c51444f833.mp4","cover_url":"http://47.113.179.123:8080/static/1_003198fdd79363536c7c6db516d174df06ed9993879b7eb2234269c51444f833.jpg","is_favorite":false,"title":"heusb"}],"next_time":1676987088}
    /home/go/douyinOrigin/test/printer.go:116: POST /douyin/user/register/?password=douyinTestUserA&username=douyinTestUserA HTTP/1.1
        Host: localhost:8080
        Content-Type: application/x-www-form-urlencoded

        password=douyinTestUserA&username=douyinTestUserA
    /home/go/douyinOrigin/test/printer.go:133: HTTP/1.1 200 OK 849.755µs
        Content-Length: 62
        Content-Type: application/json; charset=utf-8
        Date: Tue, 21 Feb 2023 13:44:48 GMT

        {"status_code":1,"status_msg":"User already exist","token":""}
    /home/go/douyinOrigin/test/printer.go:116: POST /douyin/user/login/?password=douyinTestUserA&username=douyinTestUserA HTTP/1.1
        Host: localhost:8080
        Content-Type: application/x-www-form-urlencoded

        password=douyinTestUserA&username=douyinTestUserA
    /home/go/douyinOrigin/test/printer.go:133: HTTP/1.1 200 OK 1.518259ms
        Content-Length: 196
        Content-Type: application/json; charset=utf-8
        Date: Tue, 21 Feb 2023 13:44:48 GMT

        {"status_code":0,"user_id":3,"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjo1LCJ1c2VyIjozLCJleHAiOjE2NzcwNzM0ODgsImlhdCI6MTY3Njk4NzA4OH0.RIQwDRo_Q11u-Av0vH6GTDCCou6qBHzVZgM7gCJjYYY"}
    /home/go/douyinOrigin/test/printer.go:116: POST /douyin/favorite/action/?action_type=1&token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjo1LCJ1c2VyIjozLCJleHAiOjE2NzcwNzM0ODgsImlhdCI6MTY3Njk4NzA4OH0.RIQwDRo_Q11u-Av0vH6GTDCCou6qBHzVZgM7gCJjYYY&video_id=2 HTTP/1.1
        Host: localhost:8080
        Content-Type: application/x-www-form-urlencoded

        action_type=1&token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjo1LCJ1c2VyIjozLCJleHAiOjE2NzcwNzM0ODgsImlhdCI6MTY3Njk4NzA4OH0.RIQwDRo_Q11u-Av0vH6GTDCCou6qBHzVZgM7gCJjYYY&video_id=2
    /home/go/douyinOrigin/test/printer.go:133: HTTP/1.1 200 OK 11.475292ms
        Content-Length: 40
        Content-Type: application/json; charset=utf-8
        Date: Tue, 21 Feb 2023 13:44:48 GMT

        {"status_code":0,"status_msg":"Success"}
    /home/go/douyinOrigin/test/printer.go:116: GET /douyin/favorite/list/?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjo1LCJ1c2VyIjozLCJleHAiOjE2NzcwNzM0ODgsImlhdCI6MTY3Njk4NzA4OH0.RIQwDRo_Q11u-Av0vH6GTDCCou6qBHzVZgM7gCJjYYY&user_id=3 HTTP/1.1
        Host: localhost:8080
        Content-Type: application/x-www-form-urlencoded

        token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjo1LCJ1c2VyIjozLCJleHAiOjE2NzcwNzM0ODgsImlhdCI6MTY3Njk4NzA4OH0.RIQwDRo_Q11u-Av0vH6GTDCCou6qBHzVZgM7gCJjYYY&user_id=3
    /home/go/douyinOrigin/test/printer.go:133: HTTP/1.1 200 OK 1.449418ms
        Content-Length: 409
        Content-Type: application/json; charset=utf-8
        Date: Tue, 21 Feb 2023 13:44:48 GMT

        {"status_code":0,"video_list":[{"id":2,"author":{"id":3,"name":"douyinTestUserA","work_count":1,"favorite_count":1},"play_url":"http://47.113.179.123:8080/static/3_3472a10a299cfd24a062a2f9415e57e3fd3fad14df59e9182571fef592e4eb5f.mp4","cover_url":"http://47.113.179.123:8080/static/3_3472a10a299cfd24a062a2f9415e57e3fd3fad14df59e9182571fef592e4eb5f.jpg","favorite_count":1,"is_favorite":false,"title":"Bear"}]}
--- PASS: TestFavorite (0.02s)
=== RUN   TestComment
    /home/go/douyinOrigin/test/printer.go:116: GET /douyin/feed/ HTTP/1.1
        Host: localhost:8080

    /home/go/douyinOrigin/test/printer.go:133: HTTP/1.1 200 OK 1.21459ms
        Content-Length: 761
        Content-Type: application/json; charset=utf-8
        Date: Tue, 21 Feb 2023 13:44:48 GMT

        {"status_code":0,"video_list":[{"id":2,"author":{"id":3,"name":"douyinTestUserA","work_count":1,"favorite_count":1},"play_url":"http://47.113.179.123:8080/static/3_3472a10a299cfd24a062a2f9415e57e3fd3fad14df59e9182571fef592e4eb5f.mp4","cover_url":"http://47.113.179.123:8080/static/3_3472a10a299cfd24a062a2f9415e57e3fd3fad14df59e9182571fef592e4eb5f.jpg","favorite_count":1,"is_favorite":false,"title":"Bear"},{"id":1,"author":{"id":1,"name":"test","work_count":1},"play_url":"http://47.113.179.123:8080/static/1_003198fdd79363536c7c6db516d174df06ed9993879b7eb2234269c51444f833.mp4","cover_url":"http://47.113.179.123:8080/static/1_003198fdd79363536c7c6db516d174df06ed9993879b7eb2234269c51444f833.jpg","is_favorite":false,"title":"heusb"}],"next_time":1676987088}
    /home/go/douyinOrigin/test/printer.go:116: POST /douyin/user/register/?password=douyinTestUserA&username=douyinTestUserA HTTP/1.1
        Host: localhost:8080
        Content-Type: application/x-www-form-urlencoded

        password=douyinTestUserA&username=douyinTestUserA
    /home/go/douyinOrigin/test/printer.go:133: HTTP/1.1 200 OK 706.034µs
        Content-Length: 62
        Content-Type: application/json; charset=utf-8
        Date: Tue, 21 Feb 2023 13:44:48 GMT

        {"status_code":1,"status_msg":"User already exist","token":""}
    /home/go/douyinOrigin/test/printer.go:116: POST /douyin/user/login/?password=douyinTestUserA&username=douyinTestUserA HTTP/1.1
        Host: localhost:8080
        Content-Type: application/x-www-form-urlencoded

        password=douyinTestUserA&username=douyinTestUserA
    /home/go/douyinOrigin/test/printer.go:133: HTTP/1.1 200 OK 777.666µs
        Content-Length: 196
        Content-Type: application/json; charset=utf-8
        Date: Tue, 21 Feb 2023 13:44:48 GMT

        {"status_code":0,"user_id":3,"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjo2LCJ1c2VyIjozLCJleHAiOjE2NzcwNzM0ODgsImlhdCI6MTY3Njk4NzA4OH0.yzpNYC7smLGSuBB4UgybhDgU08dVmnIoboY931vlfX0"}
    /home/go/douyinOrigin/test/printer.go:116: POST /douyin/comment/action/?action_type=1&comment_text=%E6%B5%8B%E8%AF%95%E8%AF%84%E8%AE%BA&token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjo2LCJ1c2VyIjozLCJleHAiOjE2NzcwNzM0ODgsImlhdCI6MTY3Njk4NzA4OH0.yzpNYC7smLGSuBB4UgybhDgU08dVmnIoboY931vlfX0&video_id=2 HTTP/1.1
        Host: localhost:8080
        Content-Type: application/x-www-form-urlencoded

        action_type=1&comment_text=%E6%B5%8B%E8%AF%95%E8%AF%84%E8%AE%BA&token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjo2LCJ1c2VyIjozLCJleHAiOjE2NzcwNzM0ODgsImlhdCI6MTY3Njk4NzA4OH0.yzpNYC7smLGSuBB4UgybhDgU08dVmnIoboY931vlfX0&video_id=2
    /home/go/douyinOrigin/test/printer.go:133: HTTP/1.1 200 OK 71.941649ms
        Content-Length: 208
        Content-Type: application/json; charset=utf-8
        Date: Tue, 21 Feb 2023 13:44:48 GMT

        {"status_code":0,"status_msg":"publish successfully","comment":{"id":1,"user":{"id":3,"name":"douyinTestUserA","work_count":1,"favorite_count":1},"content":"测试评论","create_date":"2023-02-21 13:44:48"}}
    /home/go/douyinOrigin/test/printer.go:116: GET /douyin/comment/list/?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjo2LCJ1c2VyIjozLCJleHAiOjE2NzcwNzM0ODgsImlhdCI6MTY3Njk4NzA4OH0.yzpNYC7smLGSuBB4UgybhDgU08dVmnIoboY931vlfX0&video_id=2 HTTP/1.1
        Host: localhost:8080
        Content-Type: application/x-www-form-urlencoded

        token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjo2LCJ1c2VyIjozLCJleHAiOjE2NzcwNzM0ODgsImlhdCI6MTY3Njk4NzA4OH0.yzpNYC7smLGSuBB4UgybhDgU08dVmnIoboY931vlfX0&video_id=2
    /home/go/douyinOrigin/test/printer.go:133: HTTP/1.1 200 OK 1.60325ms
        Content-Length: 179
        Content-Type: application/json; charset=utf-8
        Date: Tue, 21 Feb 2023 13:44:48 GMT

        {"status_code":0,"comment_list":[{"id":1,"user":{"id":3,"name":"douyinTestUserA","work_count":1,"favorite_count":1},"content":"测试评论","create_date":"2023-02-21 13:44:48"}]}
    /home/go/douyinOrigin/test/printer.go:116: POST /douyin/comment/action/?action_type=2&comment_id=1&token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjo2LCJ1c2VyIjozLCJleHAiOjE2NzcwNzM0ODgsImlhdCI6MTY3Njk4NzA4OH0.yzpNYC7smLGSuBB4UgybhDgU08dVmnIoboY931vlfX0&video_id=2 HTTP/1.1
        Host: localhost:8080
        Content-Type: application/x-www-form-urlencoded

        action_type=2&comment_id=1&token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjo2LCJ1c2VyIjozLCJleHAiOjE2NzcwNzM0ODgsImlhdCI6MTY3Njk4NzA4OH0.yzpNYC7smLGSuBB4UgybhDgU08dVmnIoboY931vlfX0&video_id=2
    /home/go/douyinOrigin/test/printer.go:133: HTTP/1.1 200 OK 3.21199ms
        Content-Length: 17
        Content-Type: application/json; charset=utf-8
        Date: Tue, 21 Feb 2023 13:44:48 GMT

        {"status_code":0}
--- PASS: TestComment (0.08s)
PASS
ok      douyinOrigin/test  0.115s
```

# 五、Demo 演示视频

https://github.com/1Carryon/douyinOrigin/tree/master/video/doushengDemo.mp4


# 六、项目总结与反思

> ## 1.目前仍存在的问题
>
> - 项目功能不够完善，只实现了基础和互动接口
> - 文件上传成功后若出现异常无法进行回滚
>
> 
>
> ## 2.已识别出的优化项
>
> - 循环依赖的问题。对项目架构进行重新编排，合理的运用指针
> - 数据安全问题。将隐私数据通过加盐加密的方式存入数据库中，避免明文存储。
> - 视频Feed流中大量同步调用导致性能低下。使用go协程和结构体嵌套对feed流、投稿列表进行优化
> - 同样为提高响应速度而加入协程易引发数据库请求响应数量过大导致数据库异常的问题。使用rabbitmq将需要数据库响应的请求做缓冲处理
> - 每次上传视频需要建立ftp连接和验证身份，等待时间长，用户体验不佳。利用SSH远程调用命令行的方式，将服务端和视频服务器进行解耦，同时在SSH中执行长链接以节省登录验证的时间。
>
> 
>
> ## 3.架构演进的可能性
>
> 在项目设计设计时考虑使用`Api Router`来统一暴露接口。其中最致命的就是整个 `Api Router`属于 `single point of failure`，如果在这一层出现严重代码缺陷或者流量洪峰可能引发集群宕机。
>
> 所以，在未来，可以将BFF根据业务模块拆成多个，每个BFF网关都有自己的任务。
>
> 
>
> ## 4.项目过程中的反思与总结
>
> 这是我第一次系统的写Go项目，过程自然曲折，但也乐在其中。因为时间和精力有限，仅仅只完成了基础接口和互动接口，与群里一些大佬相比不值一提。自己在go上还有许多知识需要学习。同时，通过与群里大佬的交流和自己对项目的完善，也让我学到了许多关于go的技巧与思想，受益良多。
>
> 我会继续保持这份对go的热忱，在技术的道路上不断学习进步。相信明天会更加美好！！！

# 七、对未来的展望

限于本人目前技术知识面、代码能力等问题，谨对此项目提出一下展望。错漏之处，请见谅

在未来可以将项目拆分成微服务架构，数据库拓展为集群，分库分表存储数据，模块分别布置到不同的服务器上。用dubbo-go来完成分布式，pixiu作为网关，Nacos作为注册中心，以rpc调用的方式来调用当前模块依赖其他模块的方法。做到分布式处理和解耦

<img src="pictures/展望-1678793799361.jpg" alt="展望" style="zoom: 67%;" />



