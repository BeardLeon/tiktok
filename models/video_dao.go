package models

import (
	"github.com/BeardLeon/tiktok/pkg/setting"
	"gorm.io/gorm"
	"time"
)

type Video struct {
	gorm.Model
	AuthorId int64  `gorm:"column:author_id"`
	PlayUrl  string `gorm:"column:play_url"`
	CoverUrl string `gorm:"column:cover_url"`
	Title    string `gorm:"column:title"`
}

func (Video) TableName() string {
	return "videos"
}

// GetVideosByLastTime
// 根据传入的时间来获取此时间之前的 VideoCount 条视频
func GetVideosByLastTime(lastTime time.Time) ([]Video, error) {
	videos := make([]Video, 0, setting.AppSetting.VideoCount)
	result := db.Model(&Video{}).Where("created_at<?", lastTime).Order("created_at desc").
		Limit(setting.AppSetting.VideoCount).Find(&videos)
	if result.Error != nil {
		return nil, result.Error
	}
	return videos, nil
}

func GetVideosByAuthorId(authorId int64) ([]Video, error) {
	// videos 集合
	var videos []Video

	result := db.Model(&Video{}).Where("author_id=?", authorId).Find(&videos)
	if result.Error != nil {
		return nil, result.Error
	}
	return videos, nil
}
