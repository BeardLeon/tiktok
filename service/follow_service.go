package service

import "github.com/BeardLeon/tiktok/models"

func GetFollowCountById(id int64) (int64, error) {
	return models.GetFollowCountById(id)
}

func GetFollowerCountById(id int64) (int64, error) {
	return models.GetFollowerCountById(id)
}

func GetIsFollow(id, authorId int64) (bool, error) {
	return models.GetIsFollow(id, authorId)
}
