package main

import (
	"context"
	"flag"

	"github.com/kolo7/bench-tpl/config"
	"github.com/kolo7/bench-tpl/generate"
	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/config.yaml", "the config file")

func main() {
	var c config.Config
	flag.Parse()
	conf.MustLoad(*configFile, &c)

	dbGenerator := generate.NewDBGenerator(&c)
	err := dbGenerator.Generate(context.Background())
	if err != nil {
		panic(err)
	}
}
