package controller

import (
	"github.com/BeardLeon/tiktok/pkg/util"
	"github.com/BeardLeon/tiktok/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	// 获取参数
	token, ok := c.GetQuery("token")
	if !ok {
		c.JSON(http.StatusOK, service.Response{StatusCode: 1, StatusMsg: "Missing parameter token"})
		return
	}
	stringVideoId, ok := c.GetQuery("video_id")
	if !ok {
		c.JSON(http.StatusOK, service.Response{StatusCode: 1, StatusMsg: "Missing parameter video_id"})
		return
	}
	videoId, err := strconv.ParseInt(stringVideoId, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, service.Response{StatusCode: 1, StatusMsg: "video_id error"})
		return
	}
	actionType, ok := c.GetQuery("action_type")
	if !ok {
		c.JSON(http.StatusOK, service.Response{StatusCode: 1, StatusMsg: "Missing parameter action_type"})
		return
	}

	// 这里要保证登录的用户的 token 一定可以通过 Redis 查询到
	if _, exist := userIsLogin(token); !exist {
		c.JSON(http.StatusOK, service.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	// 解析 token
	claims, err := util.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusOK, service.Response{StatusCode: 1, StatusMsg: "Parse token error"})
		return
	}

	// 查询用户 id
	user, err := service.GetUserByNameAndPassword(claims.Username, claims.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			service.Response{StatusCode: 1, StatusMsg: "service.GetUserByNameAndPassword error"})
		return
	}

	if actionType != "1" && actionType != "2" {
		c.JSON(http.StatusOK, service.Response{StatusCode: 0, StatusMsg: "actionType error"})
		return
	}
	if actionType == "1" {
		if err = service.Favorite(int64(user.ID), videoId); err != nil {
			c.JSON(http.StatusInternalServerError,
				service.Response{StatusCode: 1, StatusMsg: "service.Favorite error"})
			return
		}
	} else if actionType == "2" {
		if err = service.CancelFavorite(int64(user.ID), videoId); err != nil {
			c.JSON(http.StatusInternalServerError,
				service.Response{StatusCode: 1, StatusMsg: "service.CancelFavorite error"})
			return
		}
	}
	c.JSON(http.StatusOK, service.Response{StatusCode: 0, StatusMsg: "Successful"})
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	// 获取参数
	token, ok := c.GetQuery("token")
	if !ok {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: service.Response{StatusCode: 1, StatusMsg: "Missing parameter token"},
		})
		return
	}
	stringUserId, ok := c.GetQuery("user_id")
	if !ok {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: service.Response{StatusCode: 1, StatusMsg: "Missing parameter user_id"},
		})
		return
	}
	userId, err := strconv.ParseInt(stringUserId, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: service.Response{StatusCode: 1, StatusMsg: "user_id error"},
		})
		return
	}

	// 这里要保证登录的用户的 token 一定可以通过 Redis 查询到
	if _, exist := userIsLogin(token); !exist {
		c.JSON(http.StatusOK, service.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	// 获取用户的所有点赞视频
	videoList, err := service.GetFavoriteVideosByUserId(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, VideoListResponse{
			Response: service.Response{StatusCode: 1, StatusMsg: "service.GetFavoriteVideosByUserId error"},
		})
		return
	}

	c.JSON(http.StatusOK, VideoListResponse{
		Response: service.Response{
			StatusCode: 0,
		},
		VideoList: videoList,
	})
}
