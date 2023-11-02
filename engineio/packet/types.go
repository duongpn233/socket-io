package packet

import (
	"github.com/duongpn233/socket-io/engineio/frame"
)

type Frame struct {
	FType frame.Type
	Data  []byte
}

type Packet struct {
	FType frame.Type
	PType Type
	Data  []byte
}
