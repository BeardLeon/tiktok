package controller

import (
	"fmt"
	"github.com/BeardLeon/tiktok/pkg/util"
	"github.com/BeardLeon/tiktok/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// var userIdSequence = int64(1)

type UserLoginResponse struct {
	service.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	service.Response
	User service.User `json:"user"`
}

const maxLen int = 32

// checkUsernameAndPassword 通过请求参数获取用户名、密码、token 以及参数是否都存在且合法
func checkUsernameAndPassword(c *gin.Context) (string, string, string, bool) {
	// TODO: print error
	username, ok := c.GetQuery("username")
	if !ok {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: service.Response{
				StatusCode: 1,
				StatusMsg:  "Missing parameter username",
			},
		})
		return "", "", "", false
	}
	password, ok := c.GetQuery("password")
	if !ok {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: service.Response{
				StatusCode: 1,
				StatusMsg:  "Missing parameter password",
			},
		})
		return "", "", "", false
	}

	if len(username) > maxLen || len(password) > maxLen {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: service.Response{
				StatusCode: 1,
				StatusMsg:  "The length of username or password exceeds the limit",
			},
		})
		return "", "", "", false
	}

	token, err := util.GenerateToken(username, password)
	if err != nil {
		fmt.Println("token error")
		c.JSON(http.StatusInternalServerError, UserLoginResponse{
			Response: service.Response{StatusCode: 1},
		})
		return "", "", "", false
	}
	return username, password, token, true
}

func Register(c *gin.Context) {
	// TODO: print error
	username, password, token, next := checkUsernameAndPassword(c)
	if !next {
		return
	}

	// TODO: 内存 map 替换为 Redis
	if _, exist := userIsLogin(token); exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: service.Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
		return
	}

	exist, err := service.IsExistByName(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, UserLoginResponse{
			Response: service.Response{StatusCode: 1},
		})
		return
	}
	if exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: service.Response{StatusCode: 1, StatusMsg: "Username already exist"},
		})
		return
	}

	newUser, err := service.CreateUser(username, password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, UserLoginResponse{
			Response: service.Response{StatusCode: 1, StatusMsg: "service.CreateUser error"},
		})
		return
	}
	if newUser == nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: service.Response{StatusCode: 1, StatusMsg: "Failed to create user"},
		})
		return
	}
	// atomic.AddInt64(&userIdSequence, 1)
	err = userLogin(token, newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, UserLoginResponse{
			Response: service.Response{StatusCode: 1, StatusMsg: "User login error"},
		})
		return
	}
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: service.Response{StatusCode: 0},
		UserId:   newUser.Id,
		Token:    token,
	})
}

func Login(c *gin.Context) {
	// TODO: print error
	username, password, token, next := checkUsernameAndPassword(c)
	if !next {
		return
	}

	if _, exist := userIsLogin(token); exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: service.Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
		return
	}

	// 查询数据库是否存在
	user, err := service.GetUserByNameAndPassword(username, password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, UserLoginResponse{
			Response: service.Response{StatusCode: 1, StatusMsg: "GetUserByNameAndPassword error"},
		})
		return
	}
	if user.ID == 0 {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: service.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
		return
	}
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: service.Response{StatusCode: 0},
		UserId:   int64(user.ID),
		Token:    token,
	})
}

func UserInfo(c *gin.Context) {
	// TODO: print error
	id, ok := c.GetQuery("user_id")
	if !ok {
		c.JSON(http.StatusOK, UserResponse{
			Response: service.Response{StatusCode: 1, StatusMsg: "Missing parameter user_id"},
		})
		return
	}
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: service.Response{StatusCode: 1, StatusMsg: "user_id error"},
		})
		return
	}
	token, ok := c.GetQuery("token")
	if !ok {
		c.JSON(http.StatusOK, UserResponse{
			Response: service.Response{StatusCode: 1, StatusMsg: "Missing parameter token"},
		})
		return
	}

	if user, exist := userIsLogin(token); exist {
		if userId != user.Id {
			c.JSON(http.StatusOK, UserResponse{
				Response: service.Response{StatusCode: 1, StatusMsg: "user_id does not match token"},
			})
			return
		}
		c.JSON(http.StatusOK, UserResponse{
			Response: service.Response{StatusCode: 0},
			User:     *user,
		})
		return
	}

	user, password, err := service.GetAuthorById(-1, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, UserLoginResponse{
			Response: service.Response{StatusCode: 1},
		})
		return
	}
	if user == nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: service.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
		return
	}
	// token 验证
	claims, err := util.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: service.Response{StatusCode: 1, StatusMsg: "Token error"},
		})
		return
	}
	if claims.Username != user.Name || claims.Password != password {
		c.JSON(http.StatusOK, UserResponse{
			Response: service.Response{StatusCode: 1, StatusMsg: "Token error"},
		})
		return
	}

	c.JSON(http.StatusOK, UserResponse{
		Response: service.Response{StatusCode: 0},
		User:     *user,
	})
}
