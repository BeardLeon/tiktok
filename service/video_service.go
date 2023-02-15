package service

import (
	"github.com/BeardLeon/tiktok/controller"
	"github.com/BeardLeon/tiktok/models"
	"time"
)

// TODO: redis and local cache

func copyVideos(videos []models.Video, results []controller.Video) {
	for i, v := range videos {
		results[i].Id, results[i].PlayUrl, results[i].CoverUrl, results[i].Title =
			int64(v.ID), v.PlayUrl, v.CoverUrl, v.Title
	}
}

func GetVideosByLastTime(lastTime time.Time) ([]controller.Video, error) {
	videos, err := models.GetVideosByLastTime(lastTime)
	if err != nil {
		return nil, err
	}
	results := make([]controller.Video, len(videos))
	copyVideos(videos, results)
	for i, v := range videos {
		results[i].Author, err = GetAuthorById(int64(v.ID), v.AuthorId)
		if err != nil {
			return nil, err
		}
		results[i].FavoriteCount, err = GetFavoriteCount(int64(v.ID))
		if err != nil {
			return nil, err
		}
		results[i].CommentCount, err = GetCommentCount(int64(v.ID))
		if err != nil {
			return nil, err
		}
		results[i].IsFavorite, err = IsFavorite(int64(v.ID), int64(v.ID))
		if err != nil {
			return nil, err
		}
	}
	return results, nil
}

func GetVideosByLastTimeAndUsername(lastTime time.Time, username string) ([]controller.Video, error) {
	// TODO: 根据用户推荐视频
	return nil, nil
}
