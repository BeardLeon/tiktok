package service

import (
	"github.com/BeardLeon/tiktok/models"
	"time"
)

// TODO: redis and local cache

// copyVideos 将 model 层的 video 查询数据库拼接到 service 层的 video
func copyVideos(videos []models.Video, results []Video) error {
	for i, v := range videos {
		results[i].Id, results[i].PlayUrl, results[i].CoverUrl, results[i].Title =
			int64(v.ID), v.PlayUrl, v.CoverUrl, v.Title
	}
	for i, v := range videos {
		author, _, err := GetAuthorById(-1, v.AuthorId)
		if author == nil {
			return nil
		}
		results[i].Author = *author
		if err != nil {
			return err
		}
		results[i].FavoriteCount, err = GetFavoriteCount(int64(v.ID))
		if err != nil {
			return err
		}
		results[i].CommentCount, err = GetCommentCount(int64(v.ID))
		if err != nil {
			return err
		}
		results[i].IsFavorite, err = IsFavorite(int64(v.ID), int64(v.ID))
		if err != nil {
			return err
		}
	}
	return nil
}

// GetVideosByLastTime 根据传入的时间获取视频
func GetVideosByLastTime(lastTime time.Time) ([]Video, error) {
	videos, err := models.GetVideosByLastTime(lastTime)
	if err != nil {
		return nil, err
	}
	results := make([]Video, len(videos))
	copyVideos(videos, results)
	return results, nil
}

// GetVideosByLastTimeAndUsername 根据传入的时间和用户名获取视频
func GetVideosByLastTimeAndUsername(lastTime time.Time, username string) ([]Video, error) {
	// TODO: 根据用户推荐视频
	return GetVideosByLastTime(lastTime)
}

// GetVideosByAuthorId 获取作者 id 为 authorId 的作者的全部视频
func GetVideosByAuthorId(authorId int64) ([]Video, error) {
	videos, err := models.GetVideosByAuthorId(authorId)
	if err != nil {
		return nil, err
	}
	results := make([]Video, len(videos))
	copyVideos(videos, results)
	return results, nil
}
