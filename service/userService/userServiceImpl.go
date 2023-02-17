package userService

import (
	"crypto/hmac"
	"crypto/sha256"
	"douyinOrigin/config"
	"douyinOrigin/dao/userDao"
	"encoding/hex"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type UserServiceImpl struct {
	UserService
}

func (u UserServiceImpl) GetTableUserList() []userDao.TableUser {
	tableUsers, err := userDao.GetTableUserList()
	if err != nil {
		log.Panicln("Err:", err.Error())
		//log.Panicln("User Not Found")
		return tableUsers
	}
	log.Panicln("Query User Success")
	return tableUsers
}

// GetTableUserByUsername 根据user_name获得User对象
func (u UserServiceImpl) GetTableUserByUsername(name string) userDao.TableUser {
	tableUser, err := userDao.GetTableUserByUsername(name)
	if err != nil {
		//log.Panicln("Err:", err.Error())
		//log.Panicln("User Not Found")
		return tableUser
	}
	//log.Panicln("Query User Success")
	return tableUser
}

// GetTableUserById 根据 user_id 获得User对象
func (u UserServiceImpl) GetTableUserById(id int64) userDao.TableUser {
	tableUser, err := userDao.GetTableUserById(id)
	if err != nil {
		//log.Panicln("Err:", err.Error())
		//log.Panicln("User Not Found")
		panic(err.Error())
		return tableUser
	}
	log.Panicln("Query User Success")
	return tableUser
}

func (u UserServiceImpl) GetUserById(id int64) (User, error) {
	user := User{}
	tableUser, err := userDao.GetTableUserById(id)
	if err != nil {
		fmt.Println("User Not Found by Id")
		return user, err
	}
	user = User{
		Id:   tableUser.Id,
		Name: tableUser.Name,
	}
	return user, nil
}

// InsertTableUser 将user插入到数据表中
func (u UserServiceImpl) InsertTableUser(tableUser *userDao.TableUser) bool {
	flag := userDao.InsertTableUser(tableUser)
	if flag == false {
		//log.Panicln("插入失败")
		fmt.Println("插入失败")
		return flag
	}
	//log.Panicln("Insert Success")
	fmt.Println("Insert Success")
	return flag
}

// EnCoder 密码加密
func EnCoder(password string) string {
	hash := hmac.New(sha256.New, []byte(password)) //创建对应的sha256哈希加密算法
	sha := hex.EncodeToString(hash.Sum(nil))
	fmt.Println("EnCoder: ", sha)
	return sha
}

// GenerateToken 根据username生成一个token
func GenerateToken(username string) string {
	u := UserService.GetTableUserByUsername(new(UserServiceImpl), username)
	fmt.Printf("generatetoken: %v\n", u)
	token := NewToken(u)
	println(token)
	return token
}

// NewToken 根据信息创建token
func NewToken(u userDao.TableUser) string {
	expiresTime := time.Now().Unix() + int64(config.OneDayOfHours)
	fmt.Printf("expiresTime: %v\n", expiresTime)
	id64 := u.Id
	fmt.Printf("id: %v\n", strconv.FormatInt(id64, 10))
	claims := jwt.StandardClaims{
		Audience:  u.Name,
		ExpiresAt: expiresTime,
		Id:        strconv.FormatInt(id64, 10),
		IssuedAt:  time.Now().Unix(),
		Issuer:    "tiktok",
		NotBefore: time.Now().Unix(),
		Subject:   "token",
	}
	var jwtSecret = []byte(config.Secret)
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
