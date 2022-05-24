package conn

import (
	"fmt"
	"net"

	"github.com/edualb/ghype/internal/logger"
)

const (
	protocolHTTP = "tcp"
	addressHTTP  = ":80"
)

type Server struct {
	Protocol string
	Address  string

	logger *logger.Logger
}

func NewServer(protocol, address string) Server {
	return Server{
		Protocol: protocol,
		Address:  address,

		logger: logger.New(),
	}
}

func (s Server) Listen() error {
	protocol := protocolHTTP
	address := addressHTTP

	if s.Address != "" {
		address = s.Address
	}

	if s.Protocol != "" {
		protocol = s.Protocol
	}

	listener, err := net.Listen(protocol, address)
	if err != nil {
		return err
	}

	s.logger.Infof(fmt.Sprintf("Listening on %s\n", address))
	for {
		s.logger.Infof("Waiting for connection...\n")
		conn, err := listener.Accept()
		if err != nil {
			s.logger.Errof(fmt.Sprintf("connection error: %s\n", err.Error()))
			continue
		}
		s.logger.Infof("Connection accepeted\n")

		conn.Write([]byte("Hello world!\n"))

		s.logger.Infof("closing connection\n")
		err = conn.Close()
		if err != nil {
			s.logger.Errof(fmt.Sprintf("closing connection error: %s\n", err.Error()))
			continue
		}
		s.logger.Infof("Connection is done\n")
	}
}
