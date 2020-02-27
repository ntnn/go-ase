package transport

import (
	"encoding/binary"
	"fmt"
)

type MessageHeaderType uint8

const (
	TDS_BUF_LANG MessageHeaderType = iota + 1
	TDS_BUF_LOGIN
	TDS_BUF_RPC
	TDS_BUF_RESPONSE
	TDS_BUF_UNFMT
	TDS_BUF_ATTN
	TDS_BUF_BULK
	TDS_BUF_SETUP
	TDS_BUF_CLOSE
	TDS_BUF_ERROR
	TDS_BUF_PROTACK
	TDS_BUF_ECHO
	TDS_BUF_LOGOUT
	TDS_BUF_ENDPARAM
	TDS_BUF_NORMAL
	TDS_BUF_URGENT
	TDS_BUF_MIGRATE
	TDS_BUF_HELLO
	TDS_BUF_CMDSEQ_NORMAL
	TDS_BUF_CMDSEQ_LOGIN
	TDS_BUF_CMDSEQ_LIVENESS
	TDS_BUF_CMDSEQ_RESERVED1
	TDS_BUF_CMDSEQ_RESERVED2
)

type MessageHeaderStatus uint8

const (
	// Last buffer in a request or response
	TDS_BUFSTAT_EOM MessageHeaderStatus = 0x1
	// Acknowledgment of last receiver attention
	TDS_BUFSTAT_ATTNACK = 0x2
	// Attention request
	TDS_BUFSTAT_ATTN = 0x4
	// Event notification
	TDS_BUFSTAT_EVENT = 0x8
	// Buffer is encrypted
	TDS_BUFSTAT_SEAL = 0x10
	// Buffer is encrypted (SQL Anywhere CMDSQ protocol)
	TDS_BUFSTAT_ENCRYPT = 0x20
	// Buffer is encrypted with symmetric key for on demand command
	// encryption
	TDS_BUFSTAT_SYMENCRYPT = 0x40
)

type MessageHeader struct {
	// Message type, e.g. for login or language command
	MsgType MessageHeaderType
	// Status, e.g. encrypted or EOM
	Status MessageHeaderStatus
	// Length of package in bytes
	Length uint16
	// Channel the packet belongs to when multiplexing
	Channel uint16
	// PacketNr for ordering when multiplexing
	PacketNr uint8
	// Allowed window size before ACK is received
	Window uint8
}

func (header MessageHeader) String() string {
	return fmt.Sprintf(
		"MsgType: %d, Status: %d, Length: %d, Channel: %d, PacketNr: %d, Window: %d",
		header.MsgType, header.Status, header.Length, header.Channel, header.PacketNr, header.Window,
	)
}

const (
	MsgLength       = 512
	MsgHeaderLength = 8
	MsgBodyLength   = MsgLength - MsgHeaderLength
)

// TODO adhere to io interfaces
func (header MessageHeader) WriteBytes(bs []byte) error {
	if len(bs) != MsgHeaderLength {
		return fmt.Errorf("target buffer has unexpected length, expected 8 bytes, buffer length is %d", len(bs))
	}

	bs[0] = byte(header.MsgType)
	bs[1] = byte(header.Status)
	binary.BigEndian.PutUint16(bs[2:4], header.Length)
	binary.BigEndian.PutUint16(bs[4:6], header.Channel)
	bs[6] = byte(header.PacketNr)
	bs[7] = byte(header.Window)
	return nil
}

// TODO adhere to io interfaces
func (header *MessageHeader) ReadBytes(bs []byte) error {
	if len(bs) != MsgHeaderLength {
		return fmt.Errorf("passed buffer has unexpected length, expected 8 bytes, buffer length is %d", len(bs))
	}

	header.MsgType = MessageHeaderType(bs[0])
	header.Status = MessageHeaderStatus(bs[1])
	uvarint, _ := binary.Uvarint(bs[2:4])
	header.Length = uint16(uvarint)
	uvarint, _ = binary.Uvarint(bs[4:6])
	header.Channel = uint16(uvarint)
	header.PacketNr = uint8(bs[6])
	header.Window = uint8(bs[7])

	return nil
}
