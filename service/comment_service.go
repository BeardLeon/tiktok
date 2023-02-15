package service

import "github.com/BeardLeon/tiktok/models"

func GetCommentCount(videoId int64) (int64, error) {
	return models.GetCommentCount(videoId)
}
