package main

import (
	"flag"
	"github.com/oikomi/rtmp_server/server"
	"fmt"
	"github.com/golang/glog"
)

func init() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "false")
}

func main() {
	flag.Parse()
	s, err := server.New(":1935")
	if err != nil {
		return
	}

	go s.Accept()
	defer s.Close()

	for {
		select {
		case c := <-s.Clients():
			glog.Info(c)
		case err := <-s.Errs():
			fmt.Println(err)
		}
	}
}
