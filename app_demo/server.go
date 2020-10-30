package main

import "myInx/net"

func main() {
	s := net.NewServer()

	s.Serve()
}
