package controller

import (
	"github.com/BeardLeon/tiktok/pkg/util"
	"github.com/BeardLeon/tiktok/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type FeedResponse struct {
	service.Response
	VideoList []service.Video `json:"video_list,omitempty"`
	NextTime  int64           `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	// TODO: error
	// 1. 获取参数
	inputTime, ok := c.GetQuery("latest_time")
	var lastTime time.Time
	if !ok {
		lastTime = time.Now()
	} else {
		now, err := strconv.ParseInt(inputTime, 10, 64)
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}
		lastTime = time.Unix(now, 0)
	}
	token, ok := c.GetQuery("token")

	var videos []service.Video
	var err error

	if ok {
		claims, err := util.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}
		videos, err = service.GetVideosByLastTimeAndUsername(lastTime, claims.Username)
	} else {
		// 2. 调用 service 层返回视频列表
		videos, err = service.GetVideosByLastTime(lastTime)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, FeedResponse{
		Response:  service.Response{StatusCode: 0},
		VideoList: videos,
		NextTime:  time.Now().Unix(),
	})
}
