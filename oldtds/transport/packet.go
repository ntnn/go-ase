package transport

import (
	"fmt"
)

// Packet represents a single packet in a message.
type Packet struct {
	Header MessageHeader
	Data   []byte
}

func (packet Packet) Bytes() []byte {
	bs := make([]byte, int(packet.Header.Length))
	packet.Header.WriteBytes(bs[:MsgHeaderLength])
	copy(bs[MsgHeaderLength:], packet.Data)
	return bs
}

func (packet Packet) String() string {
	return fmt.Sprintf(
		"Type: %d, Status: %d, Length: %d, Channel: %d, PacketNr: %d, Window: %d, DataLen: %d",
		packet.Header.MsgType,
		packet.Header.Status,
		packet.Header.Length,
		packet.Header.Channel,
		packet.Header.PacketNr,
		packet.Header.Window,
		len(packet.Data),
	)
}
