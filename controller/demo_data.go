package controller

import "github.com/BeardLeon/tiktok/service"

var DemoVideos = []service.Video{
	{
		Id:      1,
		Author:  DemoUser,
		PlayUrl: "/static/video/bear.mp4",
		// PlayUrl:       "http://clips.vorwaerts-gmbh.de/big_buck_bunny.mp4",
		CoverUrl:      "/static/upload/images/9b41a38663113c675f719b34c9572f48.png",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	},
}

var DemoComments = []service.Comment{
	{
		Id:         1,
		User:       DemoUser,
		Content:    "Test Comment",
		CreateDate: "05-01",
	},
}

var DemoUser = service.User{
	Id:            1,
	Name:          "TestUser",
	FollowCount:   0,
	FollowerCount: 0,
	IsFollow:      false,
}
