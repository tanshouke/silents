package initialize

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/tanshouke/silents/01-config/viper/global"
	"os"
)

/*
公共函数：初始化配置。程序启动之初调用
*/
func InitConfig() {
	v := viper.New()
	v.SetConfigType("yaml")

	/* 写死文件和路径 */
	v = setConfigFile(v)

	/* 参数获取 */
	//v = setConfigFileByArgs(v)

	parse(v)
}

/*
指定文件
go run  .\main.go
*/
func setConfigFile(v *viper.Viper) *viper.Viper {
	v.SetConfigFile("config.yaml")
	return v
}

/*
参数配置
go run  .\main.go "./etc/config.yaml"
*/
func setConfigFileByArgs(v *viper.Viper) *viper.Viper {
	fmt.Println(os.Args[1])
	v.SetConfigFile(os.Args[1])
	return v
}

/*
反序列化（Deserialization） json->struct :Unmarshal
*/
func parse(v *viper.Viper) {
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	// Unmarshal config into struct
	if err := v.Unmarshal(&global.Cfg); err != nil {
		panic(fmt.Errorf("unable to decode config: %w", err))
	}

	//  Override with environment variables for sensitive data
	/*
		如果是容器、k8s发布，优先使用环境变量注入
	*/
	if dbPassword := os.Getenv("MYSQL_PASSWORD"); dbPassword != "" {
		global.Cfg.MySQL.Password = dbPassword
	}
	if redisPassword := os.Getenv("REDIS_PASSWORD"); redisPassword != "" {
		global.Cfg.Redis.Password = redisPassword
	}
	if jwtSecret := os.Getenv("JWT_SECRET"); jwtSecret != "" {
		global.Cfg.JWT.SigningKey = jwtSecret
	} else if global.Cfg.JWT.SigningKey == "" {
		// Generate a default secret for development only
		if global.Cfg.System.Env == "develop" {
			global.Cfg.JWT.SigningKey = "development-secret-key-change-in-production"
		}
	}

	// Watch for config file changes
	/*
		监听配置，方便动态加载配置，热加载
	*/
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.String())
		if err := v.Unmarshal(&global.Cfg); err != nil {
			fmt.Println("Error reloading config:", err)
		}
	})

	global.VP = v

	if global.Cfg.System.Env != "production" {
		fmt.Println(global.Cfg.String())
	}
}
