package routers

import (
	"github.com/BeardLeon/tiktok/controller"
	_ "github.com/BeardLeon/tiktok/docs"
	"github.com/BeardLeon/tiktok/pkg/setting"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)

	r.Static("/static", "./runtime")

	// // 方法原型
	// // func (group *RouterGroup) StaticFS(relativePath string, fs http.FileSystem) IRoutes {
	// r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	//
	// // 获取token方法
	// r.GET("/auth", api.GetAuth)
	//
	// // swagger
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//
	// // UploadImage
	// r.POST("/upload", api.UploadImage)

	// 路由分组 127.0.0.1:8000/tiktok
	// "relativePath" 中为分组路径，与文件夹无关
	apiRouter := r.Group("/douyin")
	{
		// basic apis
		apiRouter.GET("/feed/", controller.Feed)
		apiRouter.GET("/user/", controller.UserInfo)
		apiRouter.POST("/user/register/", controller.Register)
		apiRouter.POST("/user/login/", controller.Login)
		apiRouter.POST("/publish/action/", controller.Publish)
		apiRouter.GET("/publish/list/", controller.PublishList)

		// extra apis - I
		apiRouter.POST("/favorite/action/", controller.FavoriteAction)
		apiRouter.GET("/favorite/list/", controller.FavoriteList)
		apiRouter.POST("/comment/action/", controller.CommentAction)
		apiRouter.GET("/comment/list/", controller.CommentList)

		// extra apis - II
		apiRouter.POST("/relation/action/", controller.RelationAction)
		apiRouter.GET("/relation/follow/list/", controller.FollowList)
		apiRouter.GET("/relation/follower/list/", controller.FollowerList)
		apiRouter.GET("/relation/friend/list/", controller.FriendList)
		apiRouter.GET("/message/chat/", controller.MessageChat)
		apiRouter.POST("/message/action/", controller.MessageAction)
	}

	return r
}
