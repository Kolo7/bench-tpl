package config

type Config struct {
	TableConf map[string]*TableConf `json:","`
	Output    OutputConf            `json:","`
	DB        DBConf                `json:","`
	TplConf   map[string]*TplConf   `json:","`
}

type DBConf struct {
	Dsn string `json:",optional",env:"DB_DSN"`
}

type TableConf struct {
	Epoch int `json:",optional,default=100"`
}

type OutputConf struct {
	Dir    string `json:",optional,default=./output"`
	Format string `json:",optional,default=json"`
}

type TplConf struct {
	Dir    string `json:",default=."`
	Format string `json:",default=tpl"`
	Epoch  int    `json:",optional,default=100"`
}
