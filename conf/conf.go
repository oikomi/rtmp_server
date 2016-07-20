package conf

import (
	"flag"
	"github.com/BurntSushi/toml"
)

var (
	confPath string
	Conf     *Config
)

type Config struct {
	Server *Server
}

type Server struct {
	Addr string
}

func init() {
	flag.StringVar(&confPath, "conf", "./rtmp.toml", "config path")
}

func Init() error {
	if _, err := toml.DecodeFile(confPath, &Conf); err != nil {
		return err
	}
	return nil
}
