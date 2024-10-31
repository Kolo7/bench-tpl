package config

import "io/fs"

type Config struct {
	Tables    []string
	InputDir  string                `json:","`
	OutputDir string                `json:","`
	Dsn       string                `json:",optional",env:"DB_DSN"`
	EpochConf map[string]*EpochConf `json:",optional"`
	FQDN      string                `json:","`

	FS fs.FS
}

type EpochConf struct {
	Dir    string `json:",default=."`
	Format string `json:",default=tpl"`
	Epoch  int    `json:",optional,default=100"`
}

var DefaultNestConf = []*NestConf{}

type NestConf struct {
	Name      string                 `json:",optional"`
	Package   string                 `json:",optional,inherit"`
	Nest      []*NestConf            `json:",optional"`
	PkgUnique bool                   `json:",optional,default=false"`
	Override  bool                   `json:",optional,default=false"`
	ExtMap    map[string]interface{} `json:",optional"`
}
