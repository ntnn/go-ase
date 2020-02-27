package transport

// TODO remove

import (
	"bytes"
	"net"
)

type Conn struct {
	NetConn net.Conn
}

func Dial(network, address string) (*Conn, error) {
	c, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &Conn{NetConn: c}, nil
}

func (conn *Conn) Close() error {
	return conn.NetConn.Close()
}

func (conn *Conn) Rx() (*bytes.Buffer, error) {
	buf := &bytes.Buffer{}
	_, err := buf.ReadFrom(conn.NetConn)
	if err != nil {
		return nil, err
	}
	return buf, nil
}
