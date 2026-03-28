package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	testMsg  string
	testName string
	testCmd  = &cobra.Command{
		Use:   "test",
		Short: "t",
		Long:  `cli -t "content" `,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("testMsg:", testMsg)
			fmt.Println("testName:", testName)
			return nil

		},
	}
)

func init() {
	/* 自定义指令 */
	testCmd.PersistentFlags().StringVarP(&testMsg, "test", "t", "test msg", "test cli ~")
	testCmd.PersistentFlags().StringVarP(&testName, "name", "n", "test name", "test name cli ~")
	/* 注册命令 */
	RootCmd.AddCommand(testCmd)
}
