package controller

import (
	"github.com/BeardLeon/tiktok/service"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]service.User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

func userIsLogin(token string) (*service.User, bool) {
	if user, exist := usersLoginInfo[token]; exist {
		return &user, true
	}
	return nil, false
}

func userLogin(token string, user *service.User) error {
	usersLoginInfo[token] = *user
	return nil
}
