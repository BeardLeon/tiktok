package service

import (
	"github.com/BeardLeon/tiktok/models"
	"time"
)

// TODO: redis and local cache

func GetVideosByLastTime(lastTime time.Time) ([]models.Video, error) {
	return models.GetVideosByLastTime(lastTime)
}
