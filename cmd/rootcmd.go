package cmd

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/spf13/cobra"
)

var (
	dsn string

	globalF fs.FS
)

var rootCmd = &cobra.Command{
	Use:   "bench-tpl",
	Short: "mysql sql模板生成工具,还有压测数据生成功能",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, bench-tpl!")
	},
	Args: cobra.NoArgs,
}

func Execute(f fs.FS) {
	globalF = f
	rootCmd.PersistentFlags().StringVarP(&dsn, "dsn", "d", "", "数据库连接串")
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
