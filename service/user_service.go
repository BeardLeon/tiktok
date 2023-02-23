package service

import (
	"github.com/BeardLeon/tiktok/models"
)

// copyUser 将 model 层的 User 查询数据库拼接到 service 层的 User
func copyUser(modelUser *models.User, user *User) error {
	user.Id, user.Name = int64(modelUser.ID), modelUser.Name
	var err error
	user.FollowCount, err = GetFollowCountById(user.Id)
	if err != nil {
		return err
	}
	user.FollowerCount, err = GetFollowerCountById(user.Id)
	if err != nil {
		return err
	}
	return nil
}

// GetAuthorById 根据 userId, authorId 返回拼接好的 User 指针，用户的密码 Password（用于校验 token）
func GetAuthorById(userId, authorId int64) (*User, string, error) {
	user, err := models.GetUserById(authorId)
	if err != nil {
		return nil, "", err
	}
	if user == nil {
		return &User{}, "", nil
	}
	result := &User{}
	err = copyUser(user, result)
	if err != nil {
		return nil, "", err
	}
	if userId == -1 {
		result.IsFollow = false
	} else {
		result.IsFollow, err = GetIsFollow(userId, authorId)
		if err != nil {
			return &User{}, "", err
		}
	}
	return result, user.Password, nil
}

// IsExistByName 查询用户名为 name 的账号是否存在
func IsExistByName(name string) (bool, error) {
	return models.IsExistByName(name)
}

// CreateUser 创建用户名为 name，密码为 password 的账号
func CreateUser(name, password string) (*User, error) {
	user, err := models.CreateUser(name, password)
	if err != nil {
		return &User{}, err
	}
	if user == nil {
		return nil, nil
	}
	result := User{
		Id:       int64(user.ID),
		Name:     user.Name,
		IsFollow: true,
	}
	result.FollowCount, err = GetFollowCountById(int64(user.ID))
	if err != nil {
		return &User{}, err
	}
	result.FollowerCount, err = GetFollowerCountById(int64(user.ID))
	if err != nil {
		return &User{}, err
	}
	return &result, nil
}

// GetUserByNameAndPassword 查询用户名为 name，密码为 password 的用户
func GetUserByNameAndPassword(name, password string) (*User, error) {
	user, err := models.GetUserByNameAndPassword(name, password)
	if err != nil {
		return nil, err
	}
	result := &User{}
	err = copyUser(user, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
