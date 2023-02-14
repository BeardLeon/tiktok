package models

import "github.com/jinzhu/gorm"

type comment struct {
	gorm.Model
	userId      int    `gorm:"column:user_id"`
	videoId     int    `gorm:"column:video_id"`
	commentText string `gorm:"column:comment_text"`
	cancel      int    `gorm:"column:cancel"`
}

func (c comment) TableName() string {
	return "comments"
}

func GetCommentCount(videoId int) (int64, error) {
	var count int64
	result := db.Model(&comment{}).Where("video_id=?", videoId).Find(count)
	if result.Error != nil {
		return -1, result.Error
	}
	return count, nil
}
