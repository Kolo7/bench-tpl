/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/logx"
)

var outputDir string

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "下载模板到本地",
	Run: func(cmd *cobra.Command, args []string) {
		nestDirFileName, err := fs.Glob(globalF, "tpl/*.yaml")
		if err != nil {
			logx.Errorf("glob tpl files failed: %v", err)
			os.Exit(1)
		}
		if len(nestDirFileName) == 0 {
			logx.Errorf("tpl file not found")
			os.Exit(1)
		}
		tplFile := nestDirFileName[0]
		// 如果tplFile包含多级目录，则创建目录,删除前缀"tpl/"
		outputFullFilename := filepath.Join(outputDir, tplFile[4:])
		err = createDir(filepath.Dir(outputFullFilename))
		if err != nil {
			logx.Errorf("create dir failed: %v", err)
			os.Exit(1)
		}
		// 读取模板文件内容
		data, err := readBytesFromFile(tplFile)
		if err != nil {
			logx.Errorf("read tpl file failed: %v", err)
			os.Exit(1)
		}
		// 写入本地文件
		if err := writeBytesToFile(outputFullFilename, data); err != nil {
			logx.Errorf("write tpl file failed: %v", err)
			os.Exit(1)
		}
		list, err := fs.Glob(globalF, "tpl/**/*")
		if err != nil {
			logx.Errorf("glob tpl files failed: %v", err)
			os.Exit(1)
		}
		for _, tplFile := range list {
			// 如果tplFile包含多级目录，则创建目录
			outputFullFilename := filepath.Join(outputDir, tplFile[4:])
			err := createDir(filepath.Dir(outputFullFilename))
			if err != nil {
				logx.Errorf("create dir failed: %v", err)
				os.Exit(1)
			}
			// 读取模板文件内容
			data, err := readBytesFromFile(tplFile)
			if err != nil {
				logx.Errorf("read tpl file failed: %v", err)
				os.Exit(1)
			}
			// 写入本地文件
			if err := writeBytesToFile(outputFullFilename, data); err != nil {
				logx.Errorf("write tpl file failed: %v", err)
				os.Exit(1)
			}
		}

		fmt.Println("download tpl success")
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)
	downloadCmd.Flags().StringVarP(&outputDir, "output", "o", "./templates", "输出目录")
}

// 将字节数组写到本地文件
func writeBytesToFile(filename string, data []byte) error {
	// 如果文件不存在，则创建文件
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		file, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer file.Close()
	}
	// 写入文件
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}

// 创建目录，如果目录不存在则创建
func createDir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.MkdirAll(dir, 0777)
	}
	return nil
}

// 从globalF中读取指定文件内容
func readBytesFromFile(filename string) ([]byte, error) {
	f, err := globalF.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return io.ReadAll(f)
}
