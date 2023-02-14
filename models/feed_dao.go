package models

import (
	"github.com/BeardLeon/tiktok/pkg/setting"
	"github.com/jinzhu/gorm"
	"time"
)

type video struct {
	gorm.Model
	id       int    `gorm:"column:id"`
	authorId int    `gorm:"column:author_id"`
	playUrl  string `gorm:"column:play_url"`
	coverUrl string `gorm:"column:cover_url"`
	title    string `gorm:"column:title"`
}

func (v video) TableName() string {
	return "videos"
}

// GetVideosByLastTime
// 根据传入的时间来获取此时间之前的 VideoCount 条视频
func GetVideosByLastTime(lastTime time.Time) ([]Video, error) {
	videos := make([]video, setting.AppSetting.VideoCount)
	result := db.Model(&video{}).Where("created_at<?", lastTime).Order("created_at desc").
		Limit(setting.AppSetting.VideoCount).Find(&videos)
	if result.Error != nil {
		return nil, result.Error
	}
	results := make([]Video, len(videos))
	for i, v := range videos {
		var err error
		results[i].FavoriteCount, err = GetFavoriteCount(v.id)
		if err != nil {
			return nil, err
		}
		results[i].FavoriteCount, err = GetCommentCount(v.id)
		if err != nil {
			return nil, err
		}
		results[i].IsFavorite = false
	}
	return results, nil
}
