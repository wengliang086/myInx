package net

import (
	"context"
	"net"
	"sync"
)

type Connection struct {
	Conn        *net.TCPConn
	ConnID      uint32
	ctx         context.Context
	cancel      context.CancelFunc
	msgChan     chan []byte
	msgBuffChan chan []byte

	sync.RWMutex

	property map[string]interface{}
	isClosed bool
}

func NewConnection(conn *net.TCPConn, connId uint32) *Connection {
	c := &Connection{
		Conn:        conn,
		ConnID:      connId,
		msgChan:     make(chan []byte),
		msgBuffChan: make(chan []byte, d),
		property:    make(map[string]interface{}),
		isClosed:    false,
	}
}