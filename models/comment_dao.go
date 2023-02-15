package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserId      int    `gorm:"column:user_id"`
	VideoId     int    `gorm:"column:video_id"`
	CommentText string `gorm:"column:comment_text"`
	Cancel      int    `gorm:"column:cancel"`
}

func (Comment) TableName() string {
	return "comments"
}

func GetCommentCount(videoId int64) (int64, error) {
	var count int64
	result := db.Model(&Comment{}).Where("video_id=?", videoId).Count(&count)
	if result.Error != nil {
		return -1, result.Error
	}
	return count, nil
}
