package models

import "github.com/jinzhu/gorm"

type like struct {
	gorm.Model
	userId  int `gorm:"column:user_id"`
	videoId int `gorm:"column:video_id"`
	cancel  int `gorm:"column:cancel"`
}

func (l like) TableName() string {
	return "likes"
}

func GetFavoriteCount(videoId int) (int64, error) {
	var count int64
	result := db.Model(&like{}).Where("video_id=?", videoId).Count(&count)
	if result.Error != nil {
		return -1, result.Error
	}
	return count, nil
}

func IsFavorite(userId, videoId int) (bool, error) {
	var count int64
	result := db.Model(&like{}).Where("user_id=? and video_id=?", userId, videoId).Count(&count)
	if result.Error != nil {
		return false, result.Error
	}
	return count >= 1, nil
}
