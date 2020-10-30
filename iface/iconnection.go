package iface

type IConnection interface {
	Start()
	Stop()
	SendMsg(msgId uint32, data []byte)
}
