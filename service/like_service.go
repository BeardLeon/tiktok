package service

import "github.com/BeardLeon/tiktok/models"

func GetFavoriteCount(videoId int64) (int64, error) {
	return models.GetFavoriteCount(videoId)
}

func IsFavorite(userId, videoId int64) (bool, error) {
	return models.IsFavorite(userId, videoId)
}
