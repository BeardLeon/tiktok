package models

import "gorm.io/gorm"

type Like struct {
	gorm.Model
	UserId  int64 `gorm:"column:user_id"`
	VideoId int64 `gorm:"column:video_id"`
	Cancel  int64 `gorm:"column:cancel"`
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

func Favorite(userId, videoId int64) error {
	var count int64
	result := db.Model(&Like{}).Where("user_id=? and video_id=?", userId, videoId).Count(&count)
	if result.Error != nil {
		return result.Error
	}
	if count == 0 {
		like := Like{
			UserId:  userId,
			VideoId: videoId,
			Cancel:  0,
		}
		return db.Model(&Like{}).Create(&like).Error
	}
	return db.Model(&Like{}).Where("user_id=? and video_id=?", userId, videoId).Set("cancel", 0).Error
}

func CancelFavorite(userId, videoId int64) error {
	return db.Model(&Like{}).Where("user_id=? and video_id=?", userId, videoId).Set("cancel", 1).Error
}

func GetFavoriteVideosByUserId(userId int64) ([]Video, error) {
	var likes []Like
	// 找到 userid 和 cance 全部
	err := db.Model(&Like{}).Where("user_id=? and cancel=0", userId).Find(&likes)
	// 加了一个错误判断 是不是likes有问题
	if err != nil {
		return nil, err.Error
	}
	videos := make([]Video, len(likes))
	for i, like := range likes {
		// 视频 赋给 videos 一个个
		result := db.Model(&Video{}).Where("id=?", like.VideoId).First(&videos[i])

		if result.Error != nil {
			return nil, result.Error
		}
	}
	return videos, nil
}
