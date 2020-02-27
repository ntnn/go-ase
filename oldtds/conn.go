package tds

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/SAP/go-ase/libase/tds/transport"
)

// TDSConn handles a TDS-based connection.
type TDSConn struct {
	conn net.Conn

	SecEncryptedPassword bool
	SecChallengeResponse bool
	SecTrustedUser       bool
}

func Dial(network, address string) (*TDSConn, error) {
	c, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &TDSConn{conn: c}, nil
}

func (tds *TDSConn) Close() error {
	return tds.conn.Close()
}

func (tds *TDSConn) Login(config *transport.TokenLoginConfig) error {
	loginToken, err := transport.NewTokenLogin(config)
	if err != nil {
		return fmt.Errorf("failed to create login message: %v", err)
	}

}

// Send retrieves packets from the message and sends them to the server.
// The return values are the total number of bytes (including headers)
// and the total number of packets sent.
// An error is returned when sending a packet fails.
func (tds *TDSConn) Send(msg transport.Message) (int, int, error) {
	log.Printf("sending message: %s", msg)

	totalBytes := 0
	totalPackets := 0

	outfile, err := os.OpenFile("/sybase/TST/sentbytes", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return 0, 0, fmt.Errorf("Failed to open file to write bytes to: %v", err)
	}

	for packet := range msg.Packets() {
		bs := packet.Bytes()

		if len(bs) != int(packet.Header.Length) {
			log.Printf("error: packet byte length (%d) and indicated length in header (%d) do not match",
				len(bs), packet.Header.Length)
		}

		log.Printf("writing packet: %s", packet)
		log.Printf("writing bytes: %v", bs)

		_, err := outfile.Write(bs)
		if err != nil {
			return 0, 0, fmt.Errorf("failed to write bytes to outfile: %v", err)
		}

		n, err := tds.conn.Write(bs)
		// totalBytes is added up before checking the error since an
		// error could occur while sending part of a packet.
		totalBytes += n
		if err != nil {
			return totalBytes, totalPackets, fmt.Errorf("failed to write packet %d to stream: %v", packet.Header.PacketNr, err)
		}
		// totalPackets is only incremented after it was verified that
		// the packet has been sent without an error.
		totalPackets++

		log.Printf("wrote packet")
	}

	log.Printf("sent message, wrote %d bytes", totalBytes)
	return totalBytes, totalPackets, nil
}

// TODO remove
func (tds *TDSConn) Rx() (*bytes.Buffer, error) {
	buf := &bytes.Buffer{}
	_, err := buf.ReadFrom(tds.conn)
	if err != nil {
		return nil, err
	}
	return buf, nil
}
