package main

import (
	"fmt"
	"github.com/BeardLeon/tiktok/models"
	"github.com/BeardLeon/tiktok/pkg/logging"
	"github.com/BeardLeon/tiktok/pkg/setting"
	"github.com/BeardLeon/tiktok/routers/api/v1"
	"github.com/BeardLeon/tiktok/service"
	"net/http"
)

func main() {
	go service.RunMessageServer()

	setting.Setup()
	models.Setup()
	logging.Setup()

	router := v1.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()

	// "fvbock/endless"
	// endless.DefaultReadTimeOut = setting.ReadTimeout
	// endless.DefaultWriteTimeOut = setting.WriteTimeout
	// endless.DefaultMaxHeaderBytes = 1 << 20
	// endPoint := fmt.Sprintf(":%d", setting.HTTPPort)
	//
	// server := endless.NewServer(endPoint, routers.InitRouter())
	// server.BeforeBegin = func(add string) {
	//	log.Printf("Actual pid is %d", syscall.Getpid())
	// }
	//
	// err := server.ListenAndServe()
	// if err != nil {
	//	log.Printf("Server err: %v", err)
	// }

	// Shutdown
	// go func() {
	//	if err := s.ListenAndServe(); err != nil {
	//		log.Printf("Listen: %s\n", err)
	//	}
	// }()
	//
	// quit := make(chan os.Signal)
	// signal.Notify(quit, os.Interrupt)
	// <-quit
	//
	// log.Println("Shutdown Server ...")
	//
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	// if err := s.Shutdown(ctx); err != nil {
	//	log.Fatal("Server Shutdown:", err)
	// }
	//
	// log.Println("Server exiting")
}
