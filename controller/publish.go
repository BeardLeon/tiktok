package controller

import (
	"fmt"
	"github.com/BeardLeon/tiktok/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strconv"
)

type VideoListResponse struct {
	service.Response
	VideoList []service.Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")
	// TODO: token鉴权还是没做
	// if _, exist := usersLoginInfo[token]; !exist {
	//	c.JSON(http.StatusOK, service.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	//	return
	// }

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, service.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	// 获取文件名分隔符最后面的部分
	filename := filepath.Base(data.Filename)
	user := usersLoginInfo[token]
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	saveFile := filepath.Join("./public", finalName)
	if err = c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, service.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, service.Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	authorId, ok := c.GetQuery("user_id")
	if !ok {
		c.JSON(http.StatusOK, UserResponse{
			Response: service.Response{StatusCode: 1, StatusMsg: "Missing parameter user_id"},
		})
		return
	}
	var videos []service.Video
	userId, err := strconv.ParseInt(authorId, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: service.Response{StatusCode: 1, StatusMsg: "user_id error"},
		})
	}
	videos, err = service.GetVideosByAuthorId(userId)

	c.JSON(http.StatusOK, VideoListResponse{
		Response: service.Response{
			StatusCode: 0,
		},
		VideoList: videos,
	})
}
