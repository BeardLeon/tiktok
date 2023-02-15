package service

import (
	"github.com/BeardLeon/tiktok/models"
)

func GetAuthorById(id, authorId int64) (User, error) {
	user, err := models.GetUserById(id)
	if err != nil {
		return User{}, err
	}
	author := User{
		Id:   id,
		Name: user.Name,
	}
	author.FollowCount, err = GetFollowCountById(id)
	if err != nil {
		return User{}, err
	}
	author.FollowerCount, err = GetFollowerCountById(id)
	if err != nil {
		return User{}, err
	}
	author.IsFollow, err = GetIsFollow(id, authorId)
	if err != nil {
		return User{}, err
	}
	return author, nil
}
