package controller

import (
	"encoding/json"
	go_redis "github.com/BeardLeon/tiktok/pkg/go-redis"
	"github.com/BeardLeon/tiktok/service"
	"github.com/redis/go-redis/v9"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
var usersLoginInfo map[string]service.User

const maxTokenCount = 10000

func init() {
	usersLoginInfo = make(map[string]service.User)
}

func userIsLogin(token string) (*service.User, bool) {
	// 先查询内存中是否存在该 token
	if user, exist := usersLoginInfo[token]; exist {
		return &user, true
	}
	// 内存缓存未命中则查询 redis
	data, err := go_redis.Get(token)
	if err != nil {
		if err == redis.Nil {
			return nil, false
		}
		// TODO: 处理 redis 异常
		return nil, false
	}
	user := service.User{}
	err = json.Unmarshal(data, &user)
	if err != nil {
		// TODO: 处理 json 异常
		return nil, false
	}
	return &user, true
}

func userLogin(token string, user *service.User) error {
	err := go_redis.Set(token, user)
	if err != nil {
		return err
	}
	// 内存中最多存放 maxTokenCount 个 token
	if len(usersLoginInfo) < maxTokenCount {
		usersLoginInfo[token] = *user
	}
	return nil
}
