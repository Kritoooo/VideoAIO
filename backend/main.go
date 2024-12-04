package main

import (
	"log"
	"net/http"
	"videoaio/api"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 允许跨域
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})

	// 创建处理器实例
	videoHandler := api.NewVideoHandler()

	// API 路由组
	v1 := r.Group("/api/v1")
	{
		// 视频相关路由
		video := v1.Group("/video")
		{
			video.POST("/upload", videoHandler.Upload)
			video.POST("/process", videoHandler.Process)
		}

		// 健康检查
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	// 修改这里，监听所有网络接口
	log.Fatal(r.Run("0.0.0.0:8081"))
}
