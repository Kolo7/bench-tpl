package config

type Config struct {
	Tables map[string]TableConf `json:",optional"`
	DB     DBConf               `json:",optional"`
}

type DBConf struct {
	Dsn string `json:",optional",env:"DB_DSN"`
}

type TableConf struct {
	Epoch int `json:",optional,default=100"`
}
