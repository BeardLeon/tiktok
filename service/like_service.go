package service

import "github.com/BeardLeon/tiktok/models"

// GetFavoriteCount 查询视频点赞数量
func GetFavoriteCount(videoId int64) (int64, error) {
	return models.GetFavoriteCount(videoId)
}

// IsFavorite 根据 userId 和 videoId 判断用户是否点赞过视频
func IsFavorite(userId, videoId int64) (bool, error) {
	return models.IsFavorite(userId, videoId)
}

// Favorite 用户给视频点赞
func Favorite(userId, videoId int64) error {
	return models.Favorite(userId, videoId)
}

// CancelFavorite 用户给视频取消点赞，这里客户端要保证用户已经点赞过了
func CancelFavorite(userId, videoId int64) error {
	return models.CancelFavorite(userId, videoId)
}

// GetFavoriteVideosByUserId 根据 userId 查询用户的所有点赞视频
func GetFavoriteVideosByUserId(userId int64) ([]Video, error) {
	videos, err := models.GetFavoriteVideosByUserId(userId)
	if err != nil {
		return nil, err
	}
	results := make([]Video, len(videos))
	copyVideos(videos, results)
	return results, nil
}

//
