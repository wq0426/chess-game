package http

import (
	"chess/pkg/config"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	ERROR_CODE_USER_EXIST             = 10100
	ERROR_CODE_USER_ADD_FILED         = 10101
	ERROR_CODE_DEVICE_ADD_FILED       = 10102
	ERROR_CODE_DEVICE_USER_BIND_FILED = 10103
	ERROR_CODE_FILE_UPLOAD_FILED      = 10104
	ERROR_CODE_BOT_ADD_FILED          = 10105
	ERROR_CODE_BOT_UPDATE_FILED       = 10106
	ERROR_CODE_CONFIG_GET_FILED       = 10107
	ERROR_CODE_APP_STORE_FILED        = 10108
	ERROR_CODE_ASSET_ADD_FILED        = 10109

	ERROR_MESSAGE_USER_EXIST             = "用户已存在"
	ERROR_MESSAGE_USER_ADD_FILED         = "用户添加失败"
	ERROR_MESSAGE_DEVICE_ADD_FILED       = "设备添加失败"
	ERROR_MESSAGE_DEVICE_USER_BIND_FILED = "用户设备绑定失败"
	ERROR_MESSAGE_FILE_UPLOAD_FILED      = "上传文件失败"
	ERROR_MESSAGE_BOT_ADD_FILED          = "添加bot失败"
	ERROR_MESSAGE_BOT_UPDATE_FILED       = "编辑bot失败"
	ERROR_MESSAGE_CONFIG_GET_FILED       = "获取内购项配置失败"
	ERROR_MESSAGE_APP_STORE_FILED        = "获取应用商城版本失败"
	ERROR_MESSAGE_ASSET_ADD_FILED        = "初始化权益失败"

	SERVER_ENV_LOCAL = "local"
	SERVER_ENV_TEST  = "test"
	SERVER_ENV_PROD  = "prod"
)

func Run(r *gin.Engine, addr string) {

	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	env := config.ConfigInstance.GetString("env")
	if env == "test" || env == "prod" {
		certFile := config.ConfigInstance.GetString("ssl.certificate")
		keyFile := config.ConfigInstance.GetString("ssl.private_key")
		go func() {
			if err := srv.ListenAndServeTLS(certFile, keyFile); err != nil && err != http.ErrServerClosed {
				log.Fatalf("listen: %s\n", err)
			}
		}()
	} else {
		go func() {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("listen: %s\n", err)
			}
		}()
	}
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
