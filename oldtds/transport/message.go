package transport

import "fmt"

// Message describes the methods each of the message type
// implementations must satisfy.
type Message interface {
	fmt.Stringer
	// Packets returns a channel with Packets the Message consists of.
	Packets() chan Packet
}
