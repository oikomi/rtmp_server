package server

import (
	"github.com/golang/glog"
	"github.com/oikomi/rtmp_server/client"
	"net"
)

type Server struct {
	socket  net.Listener
	clients chan *client.Client
	errs    chan error
}

func New(bind string) (server *Server, err error) {
	socket, err := net.Listen("tcp", bind)
	if err != nil {
		glog.Error(err)
		return
	}
	server = &Server{
		socket:  socket,
		clients: make(chan *client.Client),
		errs:    make(chan error),
	}
	return
}

func (s *Server) Close() (err error) {
	err = s.socket.Close()
	return
}

func (s *Server) Clients() <-chan *client.Client {
	return s.clients
}

func (s *Server) Errs() <-chan error {
	return s.errs
}

func (s *Server) Accept() {
	for {
		conn, err := s.socket.Accept()
		if err != nil {
			s.errs <- err
			continue
		}
		s.clients <- client.New(conn)
	}
}
