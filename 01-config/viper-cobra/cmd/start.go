package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tanshouke/silents/01-config/viper-cobra/initialize"
)

var (
	// Used for flags.
	startFile string
	startCmd  = &cobra.Command{
		Use:   "start",
		Short: "s",
		Long:  `config`,
		RunE: func(cmd *cobra.Command, args []string) error {
			/* cli载入配置 */
			v := viper.New()
			v.SetConfigType("yaml")
			/* 载入配置*/
			v.SetConfigFile(startFile)
			/* 配置初始化 */
			initialize.InitConfig(v)
			return nil

		},
	}
)

func init() {
	/* 自定义指令 */
	startCmd.PersistentFlags().StringVarP(&startFile, "config", "f", "./etc/cobra/config.yaml", "config file")
	/* 注册命令 */
	RootCmd.AddCommand(startCmd)
}
