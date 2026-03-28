package initialize

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/tanshouke/silents/01-config/viper-cobra/global"
)

/*
反序列化（Deserialization） json->struct :Unmarshal
*/
func InitConfig(v *viper.Viper) {
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	if err := v.Unmarshal(&global.Cfg); err != nil {
		panic(fmt.Errorf("unable to decode config: %w", err))
	}

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
