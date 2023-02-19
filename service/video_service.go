package service

import (
	"github.com/BeardLeon/tiktok/models"
	"time"
)

// TODO: redis and local cache

func copyVideos(videos []models.Video, results []Video) {
	for i, v := range videos {
		results[i].Id, results[i].PlayUrl, results[i].CoverUrl, results[i].Title =
			int64(v.ID), v.PlayUrl, v.CoverUrl, v.Title
	}
}

func GetVideosByLastTime(lastTime time.Time) ([]Video, error) {
	videos, err := models.GetVideosByLastTime(lastTime)
	if err != nil {
		return nil, err
	}
	results := make([]Video, len(videos))
	copyVideos(videos, results)
	for i, v := range videos {
		author, _, err := GetAuthorById(-1, v.AuthorId)
		if author == nil {
			return nil, nil
		}
		results[i].Author = *author
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

func GetVideosByLastTimeAndUsername(lastTime time.Time, username string) ([]Video, error) {
	// TODO: 根据用户推荐视频
	return nil, nil
}

func GetVideosByAuthorId(authorId int64) ([]Video, error) {
	videos, err := models.GetVideosByAuthorId(authorId)
	if err != nil {
		return nil, err
	}
	results := make([]Video, len(videos))
	copyVideos(videos, results)
	return results, nil
}
