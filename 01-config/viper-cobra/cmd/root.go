package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tanshouke/silents/01-config/viper-cobra/initialize"
)

func Execute() error {
	return RootCmd.Execute()
}

var (
	// Used for flags.
	CfgFile string
	RootCmd = &cobra.Command{
		Use:   "cobra-cli",
		Short: "f",
		Long:  `config`,
		RunE: func(cmd *cobra.Command, args []string) error {
			/* cli载入配置 */
			v := viper.New()
			v.SetConfigType("yaml")
			v = setConfigFileByCmd(v)

			/* 配置初始化 */
			initialize.InitConfig(v)
			return nil
		},
	}
)

func setConfigFileByCmd(v *viper.Viper) *viper.Viper {
	fmt.Println("setConfigFileByCmd：", CfgFile)
	v.SetConfigFile(CfgFile)
	return v
}
func init() {
	RootCmd.PersistentFlags().StringVarP(&CfgFile, "config", "f", "./etc/cobra/config.yaml", "config file")
}
