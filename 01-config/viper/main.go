package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tanshouke/silents/01-config/viper/initialize"
)

func main() {
	initialize.InitConfig()
	httpStart()
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
