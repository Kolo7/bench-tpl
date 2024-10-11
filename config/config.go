package config

type Config struct {
	TableConf map[string]*TableConf `json:","`
	Input     InputConf             `json:","`
	Output    OutputConf            `json:","`
	DB        DBConf                `json:","`
	EpochConf map[string]*EpochConf `json:",optional"`
	FQDN      string                `json:","`
}

type DBConf struct {
	Dsn string `json:",optional",env:"DB_DSN"`
}

type TableConf struct {
	Epoch int `json:",optional,default=100"`
}

type InputConf struct {
	Dir      string `json:",optional,default=./tpl"`
	NestFile string `json:",optional,default=./tpl/nest.yaml"`
}

type OutputConf struct {
	Dir    string `json:",optional,default=./output"`
	Format string `json:",optional,default=json"`
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
