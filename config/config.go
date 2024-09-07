package config

type Config struct {
	Tables map[string]*TableConf `json:","`
	DB     DBConf                `json:","`
	Output OutputConf            `json:","`
	Input  InputConf             `json:","`
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

type InputConf struct {
	Dir    string `json:",default=."`
	Format string `json:",default=tpl"`
}
