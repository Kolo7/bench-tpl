# Bench-tpl

mysql sql模板生成工具,还有压测数据生成功能

```
Usage:
  bench-tpl [flags]
  bench-tpl [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  download    下载模板到本地
  help        Help about any command
  model       生成model 代码

Flags:
  -d, --dsn string   数据库连接串
  -h, --help         help for bench-tpl
```

### model

mysql sql模板生成工具,还有压测数据生成功能

```
生成model 代码

Usage:
  bench-tpl model [flags]

Flags:
  -f, --fqdn string         指定生成的model的包名 (default "test")
  -h, --help                help for model
  -D, --input-dir string    指定输入模板目录 (default "tpl")
  -n, --nest-file string    指定嵌套模板文件 (default "tpl/nest.yaml")
  -o, --output-dir string   指定输出目录 (default "./output")
  -t, --tables strings      指定生成的表名
```

### download

```
下载模板到本地

Usage:
  bench-tpl download [flags]

Flags:
  -h, --help            help for download
  -o, --output string   输出目录 (default "./tpl")
```
