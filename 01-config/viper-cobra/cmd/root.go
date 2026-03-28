package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func Execute() error {
	return RootCmd.Execute()
}

var (
	CfgFile string
	RootCmd = &cobra.Command{
		Use:   "cli",
		Short: "r",
		Long:  `long msg `,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("root cli test , CfgFile: ", CfgFile)
			return nil
		},
	}
)

func init() {
	/* 自定义指令 */
	RootCmd.PersistentFlags().StringVarP(&CfgFile, "root", "r", "./etc/cobra/config.yaml", "config file")
}
