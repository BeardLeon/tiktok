package go_redis

import (
	"context"
	"encoding/json"
	"github.com/BeardLeon/tiktok/pkg/setting"
	"github.com/redis/go-redis/v9"
)

var ctx context.Context
var rdb *redis.Client

func Setup() error {
	rdb = redis.NewClient(&redis.Options{
		Addr:     setting.RedisSetting.Host,
		Password: setting.RedisSetting.Password,
		DB:       0,
	})
	ctx = context.Background()
	return nil
}

func Set(key string, data interface{}) error {
	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = rdb.Do(ctx, "SET", key, value).Result()
	if err != nil {
		return err
	}

	return nil
}

func Get(key string) ([]byte, error) {
	val, err := rdb.Do(ctx, "GET", key).Result()
	if err != nil {
		return nil, err
	}

	return val.([]byte), nil
}

// func Set(key string, data interface{}, time int) error {
// 	conn := RedisConn.Get()
// 	defer conn.Close()
//
// 	value, err := json.Marshal(data)
// 	if err != nil {
// 		return err
// 	}
//
// 	_, err = conn.Do("SET", key, value)
// 	if err != nil {
// 		return err
// 	}
//
// 	_, err = conn.Do("EXPIRE", key, time)
// 	if err != nil {
// 		return err
// 	}
//
// 	return nil
// }
//
// func Exists(key string) bool {
// 	conn := RedisConn.Get()
// 	defer conn.Close()
// 	// 将命令返回转为布尔值
// 	exists, err := redis.Bool(conn.Do("EXISTS", key))
// 	if err != nil {
// 		return false
// 	}
//
// 	return exists
// }
//
// func Get(key string) ([]byte, error) {
// 	// 在连接池中获取一个活跃连接
// 	conn := RedisConn.Get()
// 	defer conn.Close()
//
// 	// 将命令返回转为 Bytes
// 	reply, err := redis.Bytes(conn.Do("GET", key))
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return reply, nil
// }
//
// func Delete(key string) (bool, error) {
// 	conn := RedisConn.Get()
// 	defer conn.Close()
//
// 	return redis.Bool(conn.Do("DEL", key))
// }
//
// func LikeDeletes(key string) error {
// 	conn := RedisConn.Get()
// 	defer conn.Close()
//
// 	// 将命令返回转为 []string
// 	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
// 	if err != nil {
// 		return err
// 	}
//
// 	for _, key := range keys {
// 		_, err = Delete(key)
// 		if err != nil {
// 			return err
// 		}
// 	}
//
// 	return nil
// }
