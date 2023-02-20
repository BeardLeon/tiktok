package service

import (
	"github.com/BeardLeon/tiktok/models"
)

// GetAuthorById 根据 userId, authorId 返回拼接好的 User 指针，用户的密码 Password（用于校验 token）
func GetAuthorById(userId, authorId int64) (*User, string, error) {
	user, err := models.GetUserById(authorId)
	if err != nil {
		return &User{}, "", err
	}
	if user == nil {
		return nil, "", nil
	}
	result := User{
		Id:   authorId,
		Name: user.Name,
	}
	result.FollowCount, err = GetFollowCountById(authorId)
	if err != nil {
		return &User{}, "", err
	}
	result.FollowerCount, err = GetFollowerCountById(authorId)
	if err != nil {
		return &User{}, "", err
	}
	if userId == -1 {
		result.IsFollow = false
	} else {
		result.IsFollow, err = GetIsFollow(userId, authorId)
		if err != nil {
			return &User{}, "", err
		}
	}
	return &result, user.Password, nil
}

func IsExistByName(name string) (bool, error) {
	return models.IsExistByName(name)
}

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

func IsExistByNameAndPassword(name, password string) (bool, int64, error) {
	return models.IsExistByNameAndPassword(name, password)
}