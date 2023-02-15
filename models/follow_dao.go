package models

import "gorm.io/gorm"

type Follow struct {
	gorm.Model
	UserId     int64 `gorm:"column:user_id"`
	FollowerId int64 `gorm:"column:follower_id"`
	Cancel     int64 `gorm:"column:cancel"`
}

func (Follow) TableName() string {
	return "follows"
}

func GetFollowCountById(id int64) (int64, error) {
	var count int64
	result := db.Model(&Follow{}).Where("follower_id=?", id).Count(&count)
	if result.Error != nil {
		return -1, result.Error
	}
	return count, nil
}

func GetFollowerCountById(id int64) (int64, error) {
	var count int64
	result := db.Model(&Follow{}).Where("user_id=?", id).Count(&count)
	if result.Error != nil {
		return -1, result.Error
	}
	return count, nil
}

func GetIsFollow(id, authorId int64) (bool, error) {
	var count int64
	result := db.Model(&Follow{}).Where("user_id=? and follower_id=?", authorId, id).Count(&count)
	if result.Error != nil {
		return false, result.Error
	}
	return count == 1, nil
}
