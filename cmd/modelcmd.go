package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/Kolo7/bench-tpl/config"
	"github.com/Kolo7/bench-tpl/generate"
	"github.com/Kolo7/bench-tpl/utils"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/logx"
)

var (
	input struct {
		dir      string
		nestFile string
	}
	output struct {
		dir string
	}
	tables []string
	fqdn   string
)

func init() {
	modelCmd.PersistentFlags().StringVarP(&input.dir, "input-dir", "D", "tpl", "指定输入模板目录")
	modelCmd.PersistentFlags().StringVarP(&input.nestFile, "nest-file", "n", "tpl/nest.yaml", "指定嵌套模板文件")
	modelCmd.PersistentFlags().StringVarP(&output.dir, "output-dir", "o", "./output", "指定输出目录")
	modelCmd.PersistentFlags().StringSliceVarP(&tables, "tables", "t", []string{}, "指定生成的表名")
	modelCmd.PersistentFlags().StringVarP(&fqdn, "fqdn", "f", "test", "指定生成的model的包名")
	rootCmd.AddCommand(modelCmd)
}

var modelCmd = &cobra.Command{
	Use:   "model",
	Short: "生成model 代码",
	Run: func(cmd *cobra.Command, args []string) {
		var c config.Config
		c.Dsn = dsn
		c.FQDN = fqdn
		c.Tables = tables
		c.InputDir = input.dir
		c.OutputDir = output.dir
		c.FS = globalF
		data, err := loadNestConfig(input.nestFile)
		if err != nil {
			logx.Errorf("read nest config file failed: %v", err)
		}
		jsonData, err := utils.YamlToJson(data)
		if err != nil {
			logx.Errorf("yaml to json failed: %v", err)
			os.Exit(1)
		}

		if err := json.Unmarshal(jsonData, &config.DefaultNestConf); err != nil {
			fmt.Printf("unmarshal nest config failed: %v", err)
			os.Exit(1)
		}
		dbGenerator := generate.NewDBGenerator(&c)
		if _, err := dbGenerator.Generate(context.Background()); err != nil {
			fmt.Printf("generate db failed: %v", err)
			os.Exit(1)
		}
	},
}

func loadNestConfig(filename string) ([]byte, error) {
	f, err := globalF.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return data, nil
}
