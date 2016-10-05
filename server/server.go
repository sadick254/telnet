package server

import (
	"fmt"
	"net"
)

type clients []net.Conn

var cls clients

// Server represents the ftp server
type Server struct {
	Port string
}

// NewServer returns a pointer to a new ftp server
func NewServer(port string) *Server {
	return &Server{
		Port: port,
	}
}

// Run starts a the ftp server
func (s *Server) Run() {
	addrs, err := net.ResolveTCPAddr("tcp", s.Port)
	if err != nil {
		panic("Nothing")
	}
	listener, err := net.ListenTCP("tcp", addrs)
	if err != nil {
		fmt.Println(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		cls = append(cls, conn)
		ip, err := net.ResolveTCPAddr("tcp", conn.LocalAddr().String())
		if err != nil {
			fmt.Println("Could not resolve your IP address")
			conn.Close()
		}
		conn.Write([]byte("Connected To " + ip.IP.String() + "\r\n"))
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	// accept inputs
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}
		for _, v := range cls {
			if v.RemoteAddr() != conn.RemoteAddr() {
				v.Write([]byte(buf[0:n]))
			}
		}
		fmt.Printf("%s", string(buf[0:n]))
	}
}
