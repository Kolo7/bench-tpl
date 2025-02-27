# Bench-tpl

mysql sql模板生成工具,还有压测数据生成功能

安装方式:

```
go install github.com/Kolo7/bench-tpl@latest
```

使用方式:

```
Usage:
  bench-tpl [flags]
  bench-tpl [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  download    下载模板到本地
  help        Help about any command
  http        发起网络请求
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

### http

```
读入批量脚本执行，支持调整并发和限速，可用于发起网络请求。

Usage:
  bench-tpl http [flags]

Flags:
  -c, --concurrency int   并发数量 (default 1)
  -h, --help              help for http
  -i, --input string      输入文件
  -t, --interval int      单协程内请求间隔 (default 1000)
```

### 模板微调

1. 使用 `bench-tpl download` 下载模板到本地
2. 修改模板
3. 使用 `bench-tpl model` 生成代码

### 模板可使用变量

| 变量名       | 含义               | 作用域         | 变量类型       |
|--------------|--------------------|----------------|----------------|
| fqdn         | 指定生成的model的包名 | 全局           | string         |
| pkgFullName  | 包全名             | 包级别         | string         |
| upperPkgName | 包名大写           | 包级别         | string         |
| lowerPkgName | 包名小写           | 包级别         | string         |
| modelPackageName | model包全名 | 包级别         | string         |
| daoPackageName | dao包全名 | 包级别         | string         |
| apiPackageName | api包全名 | 包级别         | string         |
| tableName | 表名 | 表级别         | string         |
| upperTableName | 表名大写 | 表级别         | string         |
| lowerTableName | 表名小写 | 表级别         | string         |
| tableColumnFields | 表字段列表 | 表级别         | []string       |
| tableColumnUpperFields | 表字段列表大写 | 表级别         | []string       |
| tablePrimaryKey | 表主键 | 表级别         | *db.Column     |
| tableUpperPrimaryKeyField | 表主键大写 | 表级别         | string         |
| tableColumns | 表字段列表 | 表级别         | []db.Column    |

__db.Column 结构体定义__

```go
type Column struct {
	Field      string `json:"field" db:"field"`
	Type       string `json:"type" db:"type"`
	Null       string `json:"null" db:"null"`
	Key        string `json:"key" db:"key"`
	Extra      string `json:"extra" db:"extra"`
	Comment    string `json:"comment" db:"comment"`
	GoType     string `json:"go_type" db:"go_type"`
	Lower      string `json:"lower" db:"lower"`
	Upper      string `json:"upper" db:"upper"`
}
```

__支持的函数__

| 函数签名 | 含义 |
|----------|------|
| rand(min, max int) int | 生成随机数 |
| randomLetters(length int) string | 生成随机字母 |
| randomNumbers(length int) string | 生成随机数字 |
| randomChinese(length int) string | 生成随机中文 |
| toTag(colName string) string | 转换为标签 |
| toUpperCamelCase(str string) string | 转换为大驼峰 |
| toLowerCamelCase(str string) string | 转换为小驼峰 |
| toSnakeCase(str string) string | 转换为蛇形 |
| inExcludedFields(val string) bool | 判断是否在排除字段中(Id, CreatedAt, CreateTime, UpdateTime, UpdatedAt, Deleted) |
