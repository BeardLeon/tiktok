package service

import (
	"github.com/BeardLeon/tiktok/models"
	"time"
)

// TODO: redis and local cache

func GetVideosByLastTime(lastTime time.Time) ([]models.Video, error) {
	return models.GetVideosByLastTime(lastTime)
}

func GetVideosByLastTimeAndUsername(lastTime time.Time, username string) ([]models.Video, error) {
	// TODO: 根据用户推荐视频
	return models.GetVideosByLastTime(lastTime)
}
