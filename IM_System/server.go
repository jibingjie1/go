package main

import (
	"fmt"
	"net"
)

type Server struct {
	IP   string
	Port int
}

//创建一个server的接口
func newServer(ip string, port int) *Server {
	server := Server{
		IP:   ip,
		Port: port,
	}
	return &server
}

//启动服务器的接口
func (s *Server) Start() {
	//socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		fmt.Println(err)
	}

	//close listen socket
	defer listener.Close()

	for {
		//accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
		}
		//do handler
		go s.Handler(conn)
	}
}

func (s *Server) Handler(conn net.Conn) {
	fmt.Println("连接成功")
}
