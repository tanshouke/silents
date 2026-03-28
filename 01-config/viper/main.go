package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tanshouke/silents/01-config/viper/global"
	"github.com/tanshouke/silents/01-config/viper/initialize"
	"go.uber.org/zap"
)

func main() {
	initialize.InitConfig()
	global.Log = initialize.InitLogger()

	// 注册延迟执行函数，Sync() 确保程序退出前同步日志缓冲区到磁盘
	defer func() {
		if err := global.Log.Sync(); err != nil {
			// 如果同步失败，记录错误日志（注意：此时可能无法写入磁盘，仅作为最后尝试）
			global.Log.Error("Failed to sync logger", zap.Error(err))
		}
	}()

	global.Log.Info("Server starting...", zap.String("env", global.Cfg.System.Env))
	global.Log.Info("Server starting...", zap.String("env", global.Cfg.System.Env))
	global.Log.Info("Server starting...", zap.String("env", global.Cfg.System.Env))
	global.Log.Info("Server starting...", zap.String("env", global.Cfg.System.Env))
	global.Log.Info("Server starting...", zap.String("env", global.Cfg.System.Env))
	//httpStart()
}

/* 启动一个http服务，验证配置热加载 */
func httpStart() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
