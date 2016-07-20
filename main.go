package main

import (
	"flag"
	"github.com/golang/glog"
	"github.com/oikomi/rtmp_server/conf"
	"github.com/oikomi/rtmp_server/server"
)

func init() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "false")
}

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		glog.Error(err)
		panic(err)
	}
	s, err := server.New(conf.Conf.Server.Addr)
	if err != nil {
		panic(err)
	}
	go s.Accept()
	defer s.Close()
	for {
		select {
		case c := <-s.Clients():
			glog.Info(c)
		case err := <-s.Errs():
			glog.Error(err)
		}
	}
}
