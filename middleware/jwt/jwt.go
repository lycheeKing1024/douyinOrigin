package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"douyinOrigin/config"
	"douyinOrigin/dao"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

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
