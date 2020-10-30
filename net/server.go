package net

import (
	"fmt"
	"myInx/iface"
	"net"
)

type Server struct {
	Name      string
	IpVersion string
	Ip        string
	Port      int
}

func NewServer() iface.IServer {
	s := &Server{
		Name:      "",
		IpVersion: "tcp4",
		Ip:        "",
		Port:      0,
	}
	return s
}

func (s *Server) Start() {
	fmt.Printf("[START] Server name: %s, listener at IP: %s, Port: %d is starting\n", s.Name, s.Ip, s.Port)

	go func() {
		//1 获取一个TCP的Addr
		addr, err := net.ResolveTCPAddr(s.IpVersion, fmt.Sprintf("%s:%d", s.Ip, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr err", err)
			return
		}

		//2 监听服务器地址
		listener, err := net.ListenTCP(s.IpVersion, addr)
		if err != nil {
			fmt.Println("listen", s.IpVersion, "err", err)
			return
		}

		//已经监听成功
		fmt.Println("started server  ", s.Name, " success, now listening...")

		//TODO server.go 应该有一个自动生成ID的方法
		//var cid uint32
		//cid = 0
		//3 启动server网络连接业务
		for true {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err ", err)
				continue
			}

			fmt.Println("Get conn remote addr = ", conn.RemoteAddr().String())

		}
	}()
}

func (s *Server) Stop() {

}

func (s *Server) Serve() {
	s.Start()

	// 阻塞
	select {}
}
