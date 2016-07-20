package server

import (
	"net"
	"github.com/oikomi/rtmp_server/client"
)

type Server struct {
	socket net.Listener
	clients chan *client.Client
	errs chan error
}

func New(bind string) (*Server, error) {
	socket, err := net.Listen("tcp", bind)
	if err != nil {
	return nil, err
	}

	return &Server{
			socket:  socket,
			clients: make(chan *client.Client),
			errs:    make(chan error),
		}, nil
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

