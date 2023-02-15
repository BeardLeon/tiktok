package models

import "github.com/jinzhu/gorm"

type Like struct {
	gorm.Model
	UserId  int `gorm:"column:user_id"`
	VideoId int `gorm:"column:video_id"`
	Cancel  int `gorm:"column:cancel"`
}

func (Like) TableName() string {
	return "likes"
}

func GetFavoriteCount(videoId int64) (int64, error) {
	var count int64
	result := db.Model(&Like{}).Where("video_id=?", videoId).Count(&count)
	if result.Error != nil {
		return -1, result.Error
	}
	return count, nil
}

func IsFavorite(userId, videoId int64) (bool, error) {
	var count int64
	result := db.Model(&Like{}).Where("user_id=? and video_id=?", userId, videoId).Count(&count)
	if result.Error != nil {
		return false, result.Error
	}
	return count == 1, nil
}
