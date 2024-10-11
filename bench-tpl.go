package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"

	"github.com/Kolo7/bench-tpl/config"
	"github.com/Kolo7/bench-tpl/generate"
	"github.com/Kolo7/bench-tpl/utils"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

var configFile = flag.String("f", "", "the config file")
var nestFile = flag.String("n", "", "the nest config file")

func main() {
	flag.Parse()
	var (
		c    config.Config
		err  error
		data = []byte{}
	)
	if configFile == nil || *configFile == "" {
		configData, err := fs.ReadFile(Content, "etc/config.yaml")
		if err != nil {
			fmt.Printf("read config file failed: %v", err)
			os.Exit(1)
		}
		conf.LoadFromYamlBytes(configData, &c)
	} else {
		conf.MustLoad(*configFile, &c)
	}

	if nestFile == nil || *nestFile == "" {
		data, err = fs.ReadFile(Content, "tpl/nest.yaml")
		if err != nil {
			fmt.Printf("read nest config file failed: %v", err)
			os.Exit(1)
		}
	} else {
		data, err = loadNestConfig(*nestFile)
		if err != nil {
			logx.Errorf("read nest config file failed: %v", err)
		}
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
}

func loadNestConfig(filename string) ([]byte, error) {
	f, err := os.OpenFile(filename, os.O_RDONLY, 0666)
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
