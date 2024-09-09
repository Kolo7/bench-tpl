package config

type Config struct {
	TableConf map[string]*TableConf `json:","`
	Input     InputConf             `json:","`
	Output    OutputConf            `json:","`
	DB        DBConf                `json:","`
	TplConf   map[string]*EpochConf `json:",optional"`
	NestRoot  *NestConf             `json:",optional"`
}

type DBConf struct {
	Dsn string `json:",optional",env:"DB_DSN"`
}

type TableConf struct {
	Epoch int `json:",optional,default=100"`
}

type InputConf struct {
	Dir string `json:",optional,default=./input"`
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

type NestConf struct {
	Name        string                 `json:",optional"`
	PackageName string                 `json:",optional,inherit"`
	Nest        []*NestConf            `json:",optional"`
	ExtMap      map[string]interface{} `json:",optional"`
}
