package service

import (
	"douyinOrigin/dao"
	"fmt"
	"log"
	"strconv"
)

type UserServiceImpl struct {
	LikeService
}

func (u UserServiceImpl) GetTableUserList() []dao.TableUser {
	tableUsers, err := dao.GetTableUserList()
	if err != nil {
		log.Panicln("Err:", err.Error())
		//log.Panicln("User Not Found")
		return tableUsers
	}
	log.Panicln("Query User Success")
	return tableUsers
}

// GetTableUserByUsername 根据user_name获得User对象
func (u UserServiceImpl) GetTableUserByUsername(name string) dao.TableUser {
	tableUser, err := dao.GetTableUserByUsername(name)
	if err != nil {
		//log.Panicln("Err:", err.Error())
		//log.Panicln("User Not Found")
		return tableUser
	}
	//log.Panicln("Query User Success")
	return tableUser
}

// GetTableUserById 根据 user_id 获得User对象
func (u UserServiceImpl) GetTableUserById(id int64) dao.TableUser {
	tableUser, err := dao.GetTableUserById(id)
	if err != nil {
		//log.Panicln("Err:", err.Error())
		//log.Panicln("User Not Found")
		//panic(err.Error())
		return tableUser
	}
	log.Panicln("Query User Success")
	return tableUser
}

// GetUserById 在未登录情况下根据user_id获取User对象
func (u UserServiceImpl) GetUserById(id int64) (User, error) {
	user := User{}
	tableUser, err := dao.GetTableUserById(id)
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

// GetUserByCurId 在登录情况（curid）下根据user_id获取User对象，curid判断是否点赞
func (u UserServiceImpl) GetUserByCurId(id int64, curId int64) (User, error) {
	fmt.Printf("curId：%v\n", curId)
	user := User{}
	tableUser, err := dao.GetTableUserById(id)
	if err != nil {
		fmt.Println("User Not Found by Id")
		return user, err
	}

	//关注
	//followCount, err := u.GetFollowingCnt(id)
	//if err != nil {
	//	log.Println("Err:", err.Error())
	//}
	//followerCount, err := u.GetFollowerCnt(id)
	//if err != nil {
	//	log.Println("Err:", err.Error())
	//}
	//isfollow, err := u.IsFollowing(curId, id)
	//if err != nil {
	//	log.Println("Err:", err.Error())
	//}
	//service := GetLikeService()
	//获取用户总共被点赞数量
	impl := LikeServiceImpl{}

	//service := GetLikeService() //解决循环依赖
	totalFavourite, _ := impl.TotalFavourite(id)
	//获取用户点赞视频的数量
	favouriteVideoCount, _ := impl.FavouriteVideoCount(id)

	//查询用户的作品
	data, _ := dao.GetVideosByAuthorId(id)

	user = User{
		Id:              tableUser.Id,
		Name:            tableUser.Name,
		Avatar:          tableUser.Avatar,
		Signature:       tableUser.Signature,
		BackgroundImage: tableUser.BackgroundImage,
		WorkCount:       int64(len(data)),
		TotalFavorited:  strconv.FormatInt(totalFavourite, 10),
		FavoriteCount:   favouriteVideoCount,
	}
	return user, nil
}

// InsertTableUser 将user插入到数据表中
func (u UserServiceImpl) InsertTableUser(tableUser *dao.TableUser) bool {
	flag := dao.InsertTableUser(tableUser)
	if flag == false {
		//log.Panicln("插入失败")
		fmt.Println("插入失败")
		return flag
	}
	//log.Panicln("Insert Success")
	fmt.Println("Insert Success")
	return flag
}
